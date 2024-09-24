package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"log"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetupSCRouter(serviceController *controllers.ServiceController) *gin.Engine {
	r := gin.Default()

	serviceRouteController := NewRouteServiceController(*serviceController)

	api := r.Group("/api")
	serviceRouteController.ServiceRoute(api)

	return r
}

func SetupSCController() controllers.ServiceController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.InitCasbin(&config)

	serviceController := controllers.NewServiceController(initializers.DB)
	serviceController.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	if err := serviceController.DB.AutoMigrate(
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

	return serviceController
}

func createProfile(t *testing.T, random *rand.Rand, cities []models.City, ethnos []models.Ethnos,
	profileTags []models.ProfileTag, bodyArts []models.BodyArt, bodyTypes []models.BodyType, hairColors []models.HairColor,
	intimateHairCuts []models.IntimateHairCut, accessTokenCookie *http.Cookie, profileRouter *gin.Engine,
	userID string) (CreateProfileResponse, error) {

	w := httptest.NewRecorder()

	payload := generateCreateProfileRequest(
		random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

	payload.BodyTypeID = nil

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
	createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
	createProfileReq.Header.Set("Content-Type", "application/json")

	profileRouter.ServeHTTP(w, createProfileReq)

	var profileResponse CreateProfileResponse
	err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

	assert.Equal(t, profileResponse.Status, "success")
	assert.NotNil(t, profileResponse.Data.ID)
	checkProfilesMatch(t,
		userID, payload, profileResponse, true, false, false)

	assert.Equal(t, http.StatusCreated, w.Code)

	return profileResponse, nil
}

func TestServiceRoutes(t *testing.T) {

	ac := SetupAuthController()

	//uc := SetupUCController()
	pc := SetupPCController()
	sc := SetupSCController()

	authRouter := SetupACRouter(&ac)
	//userRouter := SetupUCRouter(&uc)
	profileRouter := SetupPCRouter(&pc)
	serviceRouter := SetupSCRouter(&sc)

	profileTags := populateProfileTags(*pc.DB)

	cities := populateCities(*pc.DB)

	// filters

	bodyTypes := populateBodyTypes(*pc.DB)

	ethnos := populateEthnos(*pc.DB)

	hairColors := populateHairColors(*pc.DB)

	intimateHairCuts := populateIntimateHairCuts(*pc.DB)

	bodyArts := populateBodyArts(*pc.DB)

	//createOwnerUser(profileController.DB)

	random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))

	t.Cleanup(func() {
		utils.CleanupTestUsers(pc.DB)
		utils.DropAllTables(pc.DB)
	})

	t.Run("POST /api/services/: fail without access_token", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		client := generateUser(random, authRouter, t)

		accessTokenCookie, _ := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, user.ID.String())

		log.Printf(profile.Status)

		payload := &models.CreateServiceRequest{
			ClientUserID:        client.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
		}

		w := httptest.NewRecorder()

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createServiceReq, _ := http.NewRequest("POST", "/api/services/", bytes.NewBuffer(jsonPayload))
		createServiceReq.Header.Set("Content-Type", "application/json")

		serviceRouter.ServeHTTP(w, createServiceReq)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("POST /api/services/: success with client's access_token", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		client := generateUser(random, authRouter, t)

		clientAccessTokenCookie, _ := loginUserGetAccessToken(t, client.Password, client.TelegramUserID, authRouter)

		accessTokenCookie, _ := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, user.ID.String())

		log.Printf(profile.Status)

		payload := &models.CreateServiceRequest{
			ClientUserID:        client.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
		}

		w := httptest.NewRecorder()

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createServiceReq, _ := http.NewRequest("POST", "/api/services/", bytes.NewBuffer(jsonPayload))
		createServiceReq.AddCookie(&http.Cookie{Name: clientAccessTokenCookie.Name, Value: clientAccessTokenCookie.Value})
		createServiceReq.Header.Set("Content-Type", "application/json")

		serviceRouter.ServeHTTP(w, createServiceReq)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("POST /api/services/: success with profile author's access_token", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		client := generateUser(random, authRouter, t)

		accessTokenCookie, _ := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, user.ID.String())

		log.Printf(profile.Status)

		payload := &models.CreateServiceRequest{
			ClientUserID:        client.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
		}

		w := httptest.NewRecorder()

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createServiceReq, _ := http.NewRequest("POST", "/api/services/", bytes.NewBuffer(jsonPayload))
		createServiceReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createServiceReq.Header.Set("Content-Type", "application/json")

		serviceRouter.ServeHTTP(w, createServiceReq)

		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		getServiceReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/services/%s", profile.Data.ID.String()), bytes.NewBuffer(jsonPayload))
		getServiceReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		getServiceReq.Header.Set("Content-Type", "application/json")

		serviceRouter.ServeHTTP(w, getServiceReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var servicesResponse ServicesResponse
		err = json.Unmarshal(w.Body.Bytes(), &servicesResponse)

		assert.NoError(t, err)
		assert.Equal(t, servicesResponse.Status, "success")
		assert.True(t, servicesResponse.Length >= 1)
		assert.Equal(t, servicesResponse.Data[0].TrustedDistance, true)
		assert.True(t, servicesResponse.Data[0].DistanceBetweenUsers <= 100)

		assert.NotNil(t, servicesResponse.Data[0].ID)

		assert.Nil(t, servicesResponse.Data[0].ClientUserRating)
		assert.Nil(t, servicesResponse.Data[0].ClientUserRatingID)
		assert.Nil(t, servicesResponse.Data[0].ProfileRatingID)
		assert.Nil(t, servicesResponse.Data[0].ProfileRating)

		assert.Equal(t, client.ID, servicesResponse.Data[0].ClientUserID)
		assert.Equal(t, profile.Data.ID, servicesResponse.Data[0].ProfileID)

	})

}
