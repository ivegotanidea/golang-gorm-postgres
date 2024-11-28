package routes

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/controllers"
	"github.com/ivegotanidea/golang-gorm-postgres/initializers"
	"github.com/ivegotanidea/golang-gorm-postgres/models"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"testing"
)

func SetupICRouter(imageController *controllers.ImageController) *gin.Engine {
	r := gin.Default()

	imageRouteController := NewRouteImageController(*imageController)

	api := r.Group("/api")
	imageRouteController.ImageRoute(api)

	return r
}

func SetupICController() controllers.ImageController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.InitCasbin(&config)

	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	imageController := controllers.NewImageController(
		initializers.DB,
		config.ImgProxyBaseUrl,
		config.S3Endpoint,
		config.ImgProxySigningHexKey,
		config.ImgProxySigningSaltHex,
		models.S3Config{
			AccessKey:    config.S3AccessKey,
			AccessSecret: config.S3AccessSecret,
			Bucket:       config.S3Bucket,
			Region:       config.S3Region,
			Endpoint:     config.S3Endpoint,
		},
		config.ProcessingGoroutinesCount)

	if err := imageController.DB.AutoMigrate(
		&models.HairColor{},
		&models.IntimateHairCut{},
		&models.Ethnos{},
		&models.BodyType{},
		&models.ProfileBodyArt{},
		&models.BodyArt{},
		&models.City{},
		&models.User{},
		&models.Profile{},
		&models.Service{},
		&models.Photo{},
		&models.ProfileOption{},
		&models.UserRating{},
		&models.ProfileRating{},
		&models.ProfileTag{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return imageController
}

// Helper function to generate a simple in-memory image
func createTestImage(format string) (io.Reader, string, error) {
	// Create a simple 100x100 image with a red background
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	red := color.RGBA{255, 0, 0, 255}
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, red)
		}
	}

	buf := new(bytes.Buffer)
	var ext string

	switch format {
	case "png":
		if err := png.Encode(buf, img); err != nil {
			return nil, "", err
		}
		ext = ".png"
	case "jpeg":
		if err := jpeg.Encode(buf, img, nil); err != nil {
			return nil, "", err
		}
		ext = ".jpg"
	default:
		return nil, "", fmt.Errorf("unsupported format: %s", format)
	}

	return buf, ext, nil
}

func TestImageRoutes(t *testing.T) {

	//ac := SetupAuthController()
	//
	////uc := SetupUCController()
	//pc := SetupPCController()
	//
	//ic := SetupICController()

	//authRouter := SetupACRouter(&ac)
	////userRouter := SetupUCRouter(&uc)
	//profileRouter := SetupPCRouter(&pc)
	//imageRouter := SetupICRouter(&ic)
	//
	//profileTags := populateProfileTags(*pc.DB)
	//
	//cities := populateCities(*pc.DB)
	//
	//// filters
	//
	//bodyTypes := populateBodyTypes(*pc.DB)
	//
	//ethnos := filterEthnosBySex(populateEthnos(*pc.DB), "female")
	//
	//hairColors := populateHairColors(*pc.DB)
	//
	//intimateHairCuts := populateIntimateHairCuts(*pc.DB)
	//
	//bodyArts := populateBodyArts(*pc.DB)
	//
	////createOwnerUser(imageController.DB)
	//
	//random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))
	//
	//t.Cleanup(func() {
	//	//utils.CleanupTestUsers(pc.DB)
	//	//utils.DropAllTables(pc.DB)
	//})

	// should be e2e tests running inside of docker
	t.Run("POST /api/images/: success with access_token", func(t *testing.T) {
		//user := generateUser(random, authRouter, t, "")
		//
		//accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		//w := httptest.NewRecorder()
		//
		//payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)
		//
		//payload.Photos = nil
		//payload.BodyTypeID = nil
		//
		//jsonPayload, err := json.Marshal(payload)
		//if err != nil {
		//	fmt.Println("Error marshaling payload:", err)
		//	return
		//}
		//
		//createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		//createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		//createProfileReq.Header.Set("Content-Type", "application/json")
		//
		//profileRouter.ServeHTTP(w, createProfileReq)
		//
		//var profileResponse CreateProfileResponse
		//err = json.Unmarshal(w.Body.Bytes(), &profileResponse)
		//
		//assert.Equal(t, profileResponse.Status, "success")
		//assert.NotNil(t, profileResponse.Data.ID)
		//checkProfilesMatch(t,
		//	user.ID.String(), payload, profileResponse, true, false, false)
		//
		//assert.Equal(t, http.StatusCreated, w.Code)
		//
		//image1, _, err := createTestImage("png")
		//assert.NoError(t, err, "Failed to create test image1")
		//
		//image2, _, err := createTestImage("jpeg")
		//assert.NoError(t, err, "Failed to create test image2")
		//
		//image3, _, err := createTestImage("png")
		//assert.NoError(t, err, "Failed to create test image3")
		//
		//// Step 4: Create a multipart form
		//var requestBody bytes.Buffer
		//writer := multipart.NewWriter(&requestBody)
		//
		//err = writer.WriteField("profileID", profileResponse.Data.ID.String())
		//assert.NoError(t, err, "Failed to write profileID field")
		//
		//// Add images to the form
		//images := []struct {
		//	Name     string
		//	Reader   io.Reader
		//	Filename string
		//}{
		//	{"images", image1, "image1.png"},
		//	{"images", image2, "image2.jpg"},
		//	{"images", image3, "image3.png"},
		//}
		//
		//for _, img := range images {
		//	part, err := writer.CreateFormFile(img.Name, img.Filename)
		//	assert.NoError(t, err, "Failed to create form file for %s", img.Filename)
		//	_, err = io.Copy(part, img.Reader)
		//	assert.NoError(t, err, "Failed to copy image data for %s", img.Filename)
		//}
		//
		//writer.Close()
		//
		//uploadReq, err := http.NewRequest("POST", "/api/images", &requestBody)
		//assert.NoError(t, err, "Failed to create POST request for image upload")
		//
		//uploadReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		//uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
		//
		//w = httptest.NewRecorder()
		//imageRouter.ServeHTTP(w, uploadReq)
		//
		//assert.Equal(t, http.StatusCreated, w.Code, "Expected status code 200 OK for image upload")
		//
		//var successResp models.SuccessResponse[models.PhotoResponse]
		//err = json.Unmarshal(w.Body.Bytes(), &successResp)
		//assert.NoError(t, err, "Failed to unmarshal image upload response")
		//
		//assert.Equal(t, "success", successResp.Status, "Expected status to be 'success'")
		//
		//getMyProfilesReq, _ := http.NewRequest("GET", "/api/profiles/my", nil)
		//getMyProfilesReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		//getMyProfilesReq.Header.Set("Content-Type", "application/json")
		//
		//w = httptest.NewRecorder()
		//profileRouter.ServeHTTP(w, getMyProfilesReq)
		//
		//var profilesResponse ProfilesResponse
		//err = json.Unmarshal(w.Body.Bytes(), &profilesResponse)
		//
		//for i, photoResp := range profilesResponse.Data[0].Photos {
		//	// Check that URLs are not empty
		//	assert.NotEmpty(t, photoResp.URL, fmt.Sprintf("Photo %d main URL should not be empty", i+1))
		//	assert.NotEmpty(t, photoResp.PreviewURL, fmt.Sprintf("Photo %d preview URL should not be empty", i+1))
		//	assert.NotEmpty(t, photoResp.PhrURL, fmt.Sprintf("Photo %d photorama URL should not be empty", i+1))
		//
		//	// Optionally, verify the URL structure
		//	expectedMainSuffix := ".webp"
		//	expectedPreviewSuffix := "_pr.webp"
		//	expectedPhotoramaSuffix := "_phr.webp"
		//
		//	assert.True(t, strings.HasSuffix(photoResp.URL, expectedMainSuffix), "Photo %d main URL should end with %s", i+1, expectedMainSuffix)
		//	assert.True(t, strings.HasSuffix(photoResp.PreviewURL, expectedPreviewSuffix), "Photo %d preview URL should end with %s", i+1, expectedPreviewSuffix)
		//	assert.True(t, strings.HasSuffix(photoResp.PhrURL, expectedPhotoramaSuffix), "Photo %d photorama URL should end with %s", i+1, expectedPhotoramaSuffix)
		//}
	})
}
