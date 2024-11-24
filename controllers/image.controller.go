package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type ImageController struct {
	DB          *gorm.DB
	baseURL     string
	signingKey  []byte
	signingSalt []byte
	storage     afero.Fs
	bucket      string

	maxWidth  int
	maxHeight int
}

func NewImageController(
	DB *gorm.DB,
	storage afero.Fs,
	baseUrl string,
	signingKeyHex string,
	signingSaltHex string,
	bucket string,
	maxWidth int,
	maxHeight int) ImageController {

	key, err := hex.DecodeString(signingKeyHex)

	if err != nil {
		log.Fatal(err, "Key expected to be hex-encoded string")
	}

	salt, err := hex.DecodeString(signingSaltHex)

	if err != nil {
		log.Fatal(err, "Salt expected to be hex-encoded string")
	}

	return ImageController{
		DB:          DB,
		baseURL:     baseUrl,
		signingKey:  key,
		signingSalt: salt,
		storage:     storage,
		bucket:      bucket,
		maxWidth:    maxWidth,
		maxHeight:   maxHeight,
	}
}

func (ic *ImageController) store(key string, image io.Reader) error {
	file, err := ic.storage.OpenFile(key, os.O_WRONLY, 0777)
	if err != nil {
		return errors.Wrap(err, "failed to open image")
	}
	defer file.Close()

	_, err = io.Copy(file, image)
	if err != nil {
		return errors.Wrap(err, "failed to store image")
	}

	return file.Close()
}

func (ic *ImageController) sign(path string) string {
	mac := hmac.New(sha256.New, ic.signingKey)
	mac.Write(ic.signingSalt)
	mac.Write([]byte(path))

	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func (ic *ImageController) generateURL(
	imgURI string,
	resize string,
	width int,
	height int,
	enlarge int,
	gravity string,
	extension string,
) string {
	imgURI = base64.RawURLEncoding.EncodeToString([]byte(imgURI))
	path := fmt.Sprintf("/rs:%s:%d:%d:%d/g:%s/%s.%s", resize, width, height, enlarge, gravity, imgURI, extension)

	return fmt.Sprintf("%s/%s%s", ic.baseURL, ic.sign(path), path)
}

func isValidResizeOption(option string) bool {
	validOptions := map[string]bool{
		"fill": true,
		"fit":  true,
		"crop": true,
		// Add other valid options
	}
	return validOptions[option]
}

func isValidGravityOption(option string) bool {
	validOptions := map[string]bool{
		"no":     true,
		"center": true,
		// Add other valid options
	}
	return validOptions[option]
}

func isValidExtension(ext string) bool {
	validExtensions := map[string]bool{
		"png":  true,
		"jpg":  true,
		"jpeg": true,
		// Add other valid extensions
	}
	return validExtensions[ext]
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

	resize := ctx.DefaultPostForm("resize", "fill")
	if !isValidResizeOption(resize) {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid resize option"})
		return
	}

	width, err := strconv.Atoi(ctx.DefaultPostForm("width", "300"))
	if err != nil || width <= 0 || width > ic.maxWidth {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid width"})
		return
	}

	height, err := strconv.Atoi(ctx.DefaultPostForm("height", "300"))
	if err != nil || height <= 0 || height > ic.maxHeight {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid height"})
		return
	}

	gravity := ctx.DefaultPostForm("gravity", "no")
	if !isValidGravityOption(gravity) {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid gravity option"})
		return
	}

	enlarge, err := strconv.Atoi(ctx.DefaultPostForm("enlarge", "1"))
	if err != nil || (enlarge != 0 && enlarge != 1) {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid enlarge value"})
		return
	}

	extension := ctx.DefaultPostForm("extension", "png")
	if !isValidExtension(extension) {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid extension"})
		return
	}

	now := time.Now()
	newImage := Photo{
		ProfileID: uuid.MustParse(profileID),
		Extension: path.Ext(fileHeader.Filename),
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

	imgKey := newImage.ID.String() + newImage.Extension

	// Store the image
	err = ic.store(imgKey, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to store image: %v", err)})
		return
	}

	imgURL := ic.generateURL(fmt.Sprintf("s3://%s/%s", ic.bucket, imgKey), resize, width, height, enlarge, gravity, extension)

	imageResponse := &ImageResponse{
		URL: imgURL,
	}

	if err := tx.Save(&newImage).Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to commit transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse[*ImageResponse]{Status: "success", Data: imageResponse})
}
