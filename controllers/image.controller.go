package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	s3 "github.com/fclairamb/afero-s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type ImageController struct {
	DB              *gorm.DB
	imgProxyBaseUrl string
	cdnBaseURL      string
	signingKey      []byte
	signingSalt     []byte
	storage         afero.Fs
	config          S3Config

	processingGoroutinesCount int
}

func NewImageController(
	DB *gorm.DB,
	baseUrl string,
	cdnBaseURL string,
	signingKeyHex string,
	signingSaltHex string,
	config S3Config,
	processingGoroutinesCount int) ImageController {

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(config.Region),
		Endpoint:         aws.String(config.Endpoint),
		Credentials:      credentials.NewStaticCredentials(config.AccessKey, config.AccessSecret, ""),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		panic(err)
	}

	storage := s3.NewFs(config.Bucket, sess)

	key, err := hex.DecodeString(signingKeyHex)

	if err != nil {
		log.Fatal(err, "Key expected to be hex-encoded string")
	}

	salt, err := hex.DecodeString(signingSaltHex)

	if err != nil {
		log.Fatal(err, "Salt expected to be hex-encoded string")
	}

	return ImageController{
		DB:                        DB,
		imgProxyBaseUrl:           baseUrl,
		cdnBaseURL:                cdnBaseURL,
		signingKey:                key,
		signingSalt:               salt,
		storage:                   storage,
		config:                    config,
		processingGoroutinesCount: processingGoroutinesCount,
	}
}

func (ic *ImageController) store(key string, image io.Reader) error {

	file, err := ic.storage.OpenFile(key, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return errors.Wrap(err, "failed to open image")
	}
	defer file.Close()

	_, err = io.Copy(file, image)
	if err != nil {
		return errors.Wrap(err, "failed to store image")
	}

	return nil
}

func (ic *ImageController) sign(path string) string {
	mac := hmac.New(sha256.New, ic.signingKey)
	mac.Write(append(ic.signingSalt, []byte(path)...))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func (ic *ImageController) delete(imgKey string) error {
	err := ic.storage.Remove(imgKey)
	if err != nil {
		return fmt.Errorf("failed to delete image from storage: %w", err)
	}
	return nil
}

func (ic *ImageController) generateImgProxyUrl(
	imgURI string,
	resize string,
	width int,
	height int,
	gravity string,
	extension string,
	quality int,
) string {
	//imgURIEncoded := base64.RawURLEncoding.EncodeToString([]byte(imgURI))

	path := fmt.Sprintf("/rs:%s:%d:%d:0/g:%s/f:%s/watermark:1:ce:0:0:0.3/q:%d/plain/%s", resize, width, height, gravity, extension, quality, imgURI)
	signature := ic.sign(path)
	imgproxyURL := fmt.Sprintf("%s/%s%s", ic.imgProxyBaseUrl, signature, path)

	return imgproxyURL
}

func (ic *ImageController) generateCDNURL(key string) string {
	return fmt.Sprintf("%s/%s", ic.cdnBaseURL, key)
}

// processSingleImage handles the processing of a single image
func (ic *ImageController) processSingleImage(fileHeader *multipart.FileHeader, profileID string) (*Photo, error) {
	// Open the uploaded file

	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Read the file into a buffer
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read image: %w", err)
	}

	newImageID := uuid.New()

	now := time.Now()
	newImage := Photo{
		ID:        newImageID,
		ProfileID: uuid.MustParse(profileID),
		Extension: ".webp", // Images will be stored in webp format
		CreatedAt: now,
		Disabled:  false,
		Deleted:   false,
		Approved:  false,
	}

	imgKeyBase := newImage.ID.String()

	// Store the original image temporarily
	tempImgKey := imgKeyBase + filepath.Ext(fileHeader.Filename)
	if err := ic.store(tempImgKey, bytes.NewReader(fileBytes)); err != nil {
		return nil, fmt.Errorf("failed to store original image: %w", err)
	}

	// Define the image versions to process
	imageVersions := []struct {
		Suffix string
		Width  int
		Height int
	}{
		{Suffix: "", Width: 1065, Height: 705},   // Main
		{Suffix: "_pr", Width: 336, Height: 504}, // Preview
		{Suffix: "_phr", Width: 60, Height: 60},  // Photorama
	}

	// Base source URL for imgproxy
	sourceImageURL := fmt.Sprintf("%s/%s/%s", ic.cdnBaseURL, ic.config.Bucket, tempImgKey)

	// Initialize a WaitGroup and mutex for concurrency
	var wg sync.WaitGroup
	var mu sync.Mutex
	processedURLs := make(map[string]string)
	processingError := make([]string, 0)

	for _, version := range imageVersions {
		wg.Add(1)

		go func(v struct {
			Suffix string
			Width  int
			Height int
		}) {
			defer wg.Done()

			// Generate the imgproxy URL
			imgproxyURL := ic.generateImgProxyUrl(
				sourceImageURL,
				"fill",
				v.Width,
				v.Height,
				"sm",
				"webp",
				30,
			)

			// Fetch the processed image from imgproxy
			resp, err := http.Get(imgproxyURL)
			if err != nil {
				mu.Lock()
				processingError = append(processingError, fmt.Sprintf("Failed to process image via imgproxy: %v", err))
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				bodyBytes, _ := io.ReadAll(resp.Body)
				mu.Lock()
				processingError = append(processingError, fmt.Sprintf("Imgproxy returned error: %s", string(bodyBytes)))
				mu.Unlock()
				return
			}

			// Read the transformed image data
			transformedImageData, err := io.ReadAll(resp.Body)
			if err != nil {
				mu.Lock()
				processingError = append(processingError, fmt.Sprintf("Failed to read transformed image data: %v", err))
				mu.Unlock()
				return
			}

			// Generate the filename/key for storing the transformed image
			versionFilename := imgKeyBase + v.Suffix + ".webp"

			// Store the transformed image
			if err := ic.store(versionFilename, bytes.NewReader(transformedImageData)); err != nil {
				mu.Lock()
				processingError = append(processingError, fmt.Sprintf("Failed to store transformed image: %v", err))
				mu.Unlock()
				return
			}

			// Generate the CDN URL
			cdnURL := ic.generateCDNURL(versionFilename)

			// Store the URL in the map
			mu.Lock()
			switch v.Suffix {
			case "":
				processedURLs["main"] = cdnURL
			case "_pr":
				processedURLs["preview"] = cdnURL
			case "_phr":
				processedURLs["photorama"] = cdnURL
			}
			mu.Unlock()
		}(version)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if any processing errors occurred
	if len(processingError) > 0 {
		return nil, fmt.Errorf(strings.Join(processingError, "; "))
	}

	err = ic.delete(tempImgKey)

	if err != nil {
		// Log the error but proceed
		log.Printf("Failed to delete temporary image: %v", err)
	}

	newImage.URL = processedURLs["main"]
	newImage.PreviewUrl = processedURLs["preview"]
	newImage.PhrURL = processedURLs["photorama"]

	return &newImage, nil
}

// UploadProfileImage godoc
//
//	@Summary		Uploads an image
//	@Description	Uploads an image file and returns the image URL
//	@Tags			Image
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			profileID	formData	string	true	"Profile ID"
//	@Param			resize		formData	string	false	"Resize option (default: fill) (options: fill, fit, crop)"
//	@Param			width		formData	int		false	"Width (default: 300)"
//	@Param			height		formData	int		false	"Height (default: 300)"
//	@Param			gravity		formData	string	false	"Gravity (default: no) (options: no, center, north, south, east, west)"
//	@Param			enlarge		formData	int		false	"Enlarge option (default: 1) (options: 0 or 1)"
//	@Param			extension	formData	string	false	"Image extension (default: png) (options: png, jpg, jpeg)"
//	@Param			image		formData	file	true	"Image file"
//	@Success		200			{object}	SuccessResponse[ImageResponse]
//	@Failure		400			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/images/upload [post]
func (ic *ImageController) UploadProfileImage(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to get image: %v", err)})
		return
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to open image: %v", err)})
		return
	}
	defer file.Close()

	profileID := ctx.PostForm("profileID")
	if profileID == "" {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "ProfileID is required"})
		return
	}

	// Read the file into a buffer
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to read image: %v", err)})
		return
	}

	// Generate a new ID for the image
	newImageID := uuid.New()

	now := time.Now()
	newImage := Photo{
		ID:        newImageID,
		ProfileID: uuid.MustParse(profileID),
		Extension: ".webp", // Images will be stored in webp format
		CreatedAt: now,
		Disabled:  false,
		Deleted:   false,
		Approved:  false,
	}

	tx := ic.DB.Begin()

	err = tx.Create(&newImage).Error
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to save image metadata",
		})
		return
	}

	imgKeyBase := newImage.ID.String()

	tempImgKey := imgKeyBase + "_temp" + filepath.Ext(fileHeader.Filename)

	err = ic.store(tempImgKey, bytes.NewReader(fileBytes))
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to store original image: %v", err)})
		return
	}

	// /rs:fill:1065:705:0/g:sm/watermark:1:ce:0:0:0.3/q:70/plain
	// /rs:fill:336:504:0/g:sm/watermark:1:ce:0:0:0.3/q:70/plain
	// /rs:fill:120:120:0/g:sm/watermark:1:ce:0:0:0.3/q:70/plain

	imageVersions := []struct {
		Suffix string
		Width  int
		Height int
	}{
		{Suffix: "", Width: 1065, Height: 705},   // Main image
		{Suffix: "_pr", Width: 336, Height: 504}, // Preview
		{Suffix: "_phr", Width: 60, Height: 60},  // Photorama
	}

	// Base source URL for imgproxy
	sourceImageURL := fmt.Sprintf("http://%s/%s", ic.config.Bucket, tempImgKey)

	for _, version := range imageVersions {
		// Generate the imgproxy URL
		imgproxyURL := ic.generateImgProxyUrl(
			sourceImageURL,
			"fill",
			version.Width,
			version.Height,
			"sm",
			"webp",
			70,
		)

		// Fetch the processed image from imgproxy
		resp, err := http.Get(imgproxyURL)
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to process image via imgproxy: %v", err)})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: fmt.Sprintf("Imgproxy returned error: %s", string(bodyBytes)),
			})
			return
		}

		// Read the transformed image data
		transformedImageData, err := io.ReadAll(resp.Body)
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to read transformed image data: %v", err)})
			return
		}

		// Generate the filename/key for storing the transformed image
		versionFilename := imgKeyBase + version.Suffix + ".webp"

		// Store the transformed image
		err = ic.store(versionFilename, bytes.NewReader(transformedImageData))
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to store transformed image: %v", err)})
			return
		}
	}

	err = ic.delete(tempImgKey)
	if err != nil {
		log.Printf("Failed to delete temporary image: %v", err)
	}

	newImage.URL = imgKeyBase + ".webp"
	newImage.PreviewUrl = imgKeyBase + "_pr.webp"
	newImage.PhrURL = imgKeyBase + "_phr.webp"

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to commit transaction",
		})
		return
	}

	photoResponse := &PhotoResponse{
		URL:        newImage.URL,
		PreviewURL: newImage.PreviewUrl,
		PhrURL:     newImage.PhrURL,
		Disabled:   newImage.Disabled,
		Approved:   newImage.Approved,
		Deleted:    newImage.Deleted,
	}

	ctx.JSON(http.StatusOK, SuccessResponse[*PhotoResponse]{Status: "success", Data: photoResponse})
}

// UploadProfileImages handles multiple image uploads and processing
//
//	@Summary		Uploads multiple images
//	@Description	Uploads multiple image files and returns their URLs
//	@Tags			Image
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			profileID	formData	string	true	"Profile ID"
//	@Param			images		formData	[]file	true	"Image files"
//	@Success		201			{object}	SuccessResponse[[]PhotoResponse]
//	@Failure		400			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/images [post]
func (ic *ImageController) UploadProfileImages(ctx *gin.Context) {
	// Retrieve all files from the "images" form field
	formFiles, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to parse multipart form: %v", err)})
		return
	}

	files := formFiles.File["images"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "No images uploaded"})
		return
	}

	profileID := ctx.PostForm("profileID")
	if profileID == "" {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "ProfileID is required"})
		return
	}

	// Initialize a slice to hold the responses
	var photos []Photo
	var processingErrors []string

	// Begin a database transaction
	tx := ic.DB.Begin()
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to start database transaction"})
		return
	}

	// Use a WaitGroup to handle concurrency
	var wg sync.WaitGroup
	var mu sync.Mutex // To protect shared slices

	// Limit the number of concurrent goroutines
	semaphore := make(chan struct{}, ic.processingGoroutinesCount) // Adjust the concurrency level as needed

	for _, fileHeader := range files {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot

		go func(fh *multipart.FileHeader) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release the slot

			// Process each file
			photo, err := ic.processSingleImage(fh, profileID)
			if err != nil {
				mu.Lock()
				processingErrors = append(processingErrors, fmt.Sprintf("Failed to process %s: %v", fh.Filename, err))
				mu.Unlock()
				return
			}

			mu.Lock()
			photos = append(photos, *photo)
			mu.Unlock()
		}(fileHeader)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if any errors occurred during processing
	if len(processingErrors) > 0 {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: strings.Join(processingErrors, "; "),
		})
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to commit transaction",
		})
		return
	}

	// Return the aggregated responses
	ctx.JSON(http.StatusCreated, nil)
}
