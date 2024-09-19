package routes

import (
	"bytes"
	"encoding/json"
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

	createOwnerUser(serviceController.DB)

	// Migrate the schema
	if err := serviceController.DB.AutoMigrate(&models.User{}, &models.Profile{}); err != nil {
		panic("failed to migrate database")
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

	authRouter := SetupACRouter(&ac)
	//userRouter := SetupUCRouter(&uc)
	profileRouter := SetupPCRouter(&pc)

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

	t.Run("POST /api/profiles/: success with access_token", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, _ := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, user.ID.String())

		log.Printf(profile.Status)
	})

}
