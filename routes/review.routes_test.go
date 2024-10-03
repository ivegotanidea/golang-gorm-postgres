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

func SetupRCRouter(serviceController *controllers.ServiceController) *gin.Engine {
	r := gin.Default()

	reviewController := NewRouteReviewController(*serviceController)

	api := r.Group("/api")
	reviewController.ReviewsRoute(api)

	return r
}

func SetupRCController() controllers.ServiceController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.InitCasbin(&config)

	serviceController := controllers.NewServiceController(initializers.DB, config.ReviewUpdateLimitHours)
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
		&models.UserTag{},
		&models.RatedUserTag{},
		&models.ProfileRating{},
		&models.ProfileTag{},
		&models.RatedProfileTag{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return serviceController
}

func TestReviewsRoutes(t *testing.T) {

	ac := SetupAuthController()

	//uc := SetupUCController()
	pc := SetupPCController()
	sc := SetupSCController()
	rc := SetupRCController()

	authRouter := SetupACRouter(&ac)
	//userRouter := SetupUCRouter(&uc)
	profileRouter := SetupPCRouter(&pc)
	serviceRouter := SetupSCRouter(&sc)
	reviewRouter := SetupRCRouter(&rc)

	profileTags := populateProfileTags(*pc.DB)
	userTags := populateUserTags(*pc.DB)

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

	// -> profile

	t.Run("PUT /api/reviews/host?service_id=:serviceID: client can't update his review after 48 hours after creation", func(t *testing.T) {

		profileOwner := generateUser(random, authRouter, t, "")
		clientUser := generateUser(random, authRouter, t, "")

		accessTokenCookie, _ := loginUserGetAccessToken(t, profileOwner.Password, profileOwner.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, profileOwner.ID.String())

		payload := &models.CreateServiceRequest{
			ClientUserID:        clientUser.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileOwnerID:       profileOwner.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
			ProfileRating: &models.CreateProfileRatingRequest{
				Review: "I like the service! It's very good",
				Score:  ptr(5),
				RatedProfileTags: []models.CreateRatedProfileTagRequest{
					{
						Type:  "like",
						TagID: profileTags[0].ID,
					},
					{
						Type:  "like",
						TagID: profileTags[1].ID,
					},
				},
			},
			UserRating: &models.CreateUserRatingRequest{
				Review: "I liked the client! He is very kind",
				Score:  ptr(5),
				RatedUserTags: []models.CreateRatedUserTagRequest{
					{
						Type:  "like",
						TagID: userTags[0].ID,
					},
					{
						Type:  "dislike",
						TagID: userTags[1].ID,
					},
				},
			},
		}

		service, _ := createServiceFromPayload(t, *payload, serviceRouter, accessTokenCookie)

		assert.NotNil(t, service)

		newCreatedAt := time.Now().UTC().Add(-49 * time.Hour)

		result := sc.DB.Model(&models.ProfileRating{}).
			Where("id = ?", service.Data[0].ProfileRatingID).
			UpdateColumn("created_at", newCreatedAt)

		if result.Error != nil {
			panic(result.Error.Error())
		}

		w := httptest.NewRecorder()

		updateProfileOwnerReviewReqBody := &models.CreateProfileRatingRequest{
			Review: "I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life",
			Score:  ptr(1),
			RatedProfileTags: []models.CreateRatedProfileTagRequest{
				{
					Type:  "dislike",
					TagID: profileTags[0].ID,
				},
			},
		}

		jsonPayload, err := json.Marshal(updateProfileOwnerReviewReqBody)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
		}

		updateUri := fmt.Sprintf("/api/reviews/host?service_id=%s", service.Data[0].ID.String())
		updateProfileOwnerReviewReq, _ := http.NewRequest("PUT", updateUri, bytes.NewBuffer(jsonPayload))
		updateProfileOwnerReviewReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		updateProfileOwnerReviewReq.Header.Set("Content-Type", "application/json")

		reviewRouter.ServeHTTP(w, updateProfileOwnerReviewReq)

		assert.Equal(t, http.StatusForbidden, w.Code)

	})

	t.Run("PUT /api/reviews/host?service_id=:serviceID: client can update hist review before 48 hours passed", func(t *testing.T) {

		profileOwner := generateUser(random, authRouter, t, "")
		clientUser := generateUser(random, authRouter, t, "")

		accessTokenCookie, _ := loginUserGetAccessToken(t, profileOwner.Password, profileOwner.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, profileOwner.ID.String())

		payload := &models.CreateServiceRequest{
			ClientUserID:        clientUser.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileOwnerID:       profileOwner.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
			ProfileRating: &models.CreateProfileRatingRequest{
				Review: "I like the service! It's very good",
				Score:  ptr(5),
				RatedProfileTags: []models.CreateRatedProfileTagRequest{
					{
						Type:  "like",
						TagID: profileTags[0].ID,
					},
					{
						Type:  "like",
						TagID: profileTags[1].ID,
					},
				},
			},
			UserRating: &models.CreateUserRatingRequest{
				Review: "I liked the client! He is very kind",
				Score:  ptr(5),
				RatedUserTags: []models.CreateRatedUserTagRequest{
					{
						Type:  "like",
						TagID: userTags[0].ID,
					},
					{
						Type:  "dislike",
						TagID: userTags[1].ID,
					},
				},
			},
		}

		service, _ := createServiceFromPayload(t, *payload, serviceRouter, accessTokenCookie)

		assert.NotNil(t, service)

		newCreatedAt := time.Now().UTC().Add(-24 * time.Hour)

		result := sc.DB.Model(&models.ProfileRating{}).
			Where("id = ?", service.Data[0].ProfileRatingID).
			UpdateColumn("created_at", newCreatedAt)

		if result.Error != nil {
			panic(result.Error.Error())
		}

		w := httptest.NewRecorder()

		updateProfileOwnerReviewReqBody := &models.CreateProfileRatingRequest{
			Review: "I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life",
			Score:  ptr(1),
			RatedProfileTags: []models.CreateRatedProfileTagRequest{
				{
					Type:  "dislike",
					TagID: profileTags[0].ID,
				},
			},
		}

		jsonPayload, err := json.Marshal(updateProfileOwnerReviewReqBody)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
		}

		updateUri := fmt.Sprintf("/api/reviews/host?profile_id=%s&service_id=%s", profile.Data.ID.String(), service.Data[0].ID.String())
		updateProfileOwnerReviewReq, _ := http.NewRequest("PUT", updateUri, bytes.NewBuffer(jsonPayload))
		updateProfileOwnerReviewReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		updateProfileOwnerReviewReq.Header.Set("Content-Type", "application/json")

		reviewRouter.ServeHTTP(w, updateProfileOwnerReviewReq)

		assert.Equal(t, http.StatusOK, w.Code)

		basicUser := generateUser(random, authRouter, t, "guru")
		basicUserAccessTokenCookie, _ := loginUserGetAccessToken(t, basicUser.Password, basicUser.TelegramUserID, authRouter)

		w = httptest.NewRecorder()
		getServiceReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/services/%s/service/%s", profile.Data.ID.String(), service.Data[0].ID.String()), nil)
		getServiceReq.AddCookie(&http.Cookie{Name: basicUserAccessTokenCookie.Name, Value: basicUserAccessTokenCookie.Value})
		getServiceReq.Header.Set("Content-Type", "application/json")

		serviceRouter.ServeHTTP(w, getServiceReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var servicesResponse ServiceResponse
		err = json.Unmarshal(w.Body.Bytes(), &servicesResponse)

		assert.NoError(t, err)

		assert.Equal(t, servicesResponse.Status, "success")

		assert.Equal(t, updateProfileOwnerReviewReqBody.Review, servicesResponse.Data.ProfileRating.Review)
		assert.Equal(t, updateProfileOwnerReviewReqBody.Score, servicesResponse.Data.ProfileRating.Score)

	})

	t.Run("PUT /api/reviews/host/set-visibility: fail profile owner basic-tier can't hide out client's review", func(t *testing.T) {

		profileOwner := generateUser(random, authRouter, t, "")
		clientUser := generateUser(random, authRouter, t, "")

		accessTokenCookie, _ := loginUserGetAccessToken(t, profileOwner.Password, profileOwner.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, profileOwner.ID.String())

		payload := &models.CreateServiceRequest{
			ClientUserID:        clientUser.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileOwnerID:       profileOwner.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
			ProfileRating: &models.CreateProfileRatingRequest{
				Review: "I like the service! It's very good",
				Score:  ptr(5),
				RatedProfileTags: []models.CreateRatedProfileTagRequest{
					{
						Type:  "like",
						TagID: profileTags[0].ID,
					},
					{
						Type:  "like",
						TagID: profileTags[1].ID,
					},
				},
			},
			UserRating: &models.CreateUserRatingRequest{
				Review: "I liked the client! He is very kind",
				Score:  ptr(5),
				RatedUserTags: []models.CreateRatedUserTagRequest{
					{
						Type:  "like",
						TagID: userTags[0].ID,
					},
					{
						Type:  "dislike",
						TagID: userTags[1].ID,
					},
				},
			},
		}

		service, _ := createServiceFromPayload(t, *payload, serviceRouter, accessTokenCookie)

		assert.NotNil(t, service)

		w := httptest.NewRecorder()

		updateProfileOwnerReviewReqBody := &models.SetReviewVisibilityRequest{
			Visible: boolPtr(false),
		}

		jsonPayload, err := json.Marshal(updateProfileOwnerReviewReqBody)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
		}

		updateUri := fmt.Sprintf("/api/reviews/host/set-visibility?service_id=%s", service.Data[0].ID.String())
		updateProfileOwnerReviewReq, _ := http.NewRequest("PUT", updateUri, bytes.NewBuffer(jsonPayload))
		updateProfileOwnerReviewReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		updateProfileOwnerReviewReq.Header.Set("Content-Type", "application/json")

		reviewRouter.ServeHTTP(w, updateProfileOwnerReviewReq)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("PUT /api/reviews/host/set-visibility: success profile owner expert-tier can't hide out client's review", func(t *testing.T) {

		profileOwner := generateUser(random, authRouter, t, "expert")
		clientUser := generateUser(random, authRouter, t, "")

		accessTokenCookie, _ := loginUserGetAccessToken(t, profileOwner.Password, profileOwner.TelegramUserID, authRouter)
		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, profileOwner.ID.String())

		payload := &models.CreateServiceRequest{
			ClientUserID:        clientUser.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileOwnerID:       profileOwner.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
			ProfileRating: &models.CreateProfileRatingRequest{
				Review: "I like the service! It's very good",
				Score:  ptr(5),
				RatedProfileTags: []models.CreateRatedProfileTagRequest{
					{
						Type:  "like",
						TagID: profileTags[0].ID,
					},
					{
						Type:  "like",
						TagID: profileTags[1].ID,
					},
				},
			},
			UserRating: &models.CreateUserRatingRequest{
				Review: "I liked the client! He is very kind",
				Score:  ptr(5),
				RatedUserTags: []models.CreateRatedUserTagRequest{
					{
						Type:  "like",
						TagID: userTags[0].ID,
					},
					{
						Type:  "dislike",
						TagID: userTags[1].ID,
					},
				},
			},
		}

		service, _ := createServiceFromPayload(t, *payload, serviceRouter, accessTokenCookie)

		assert.NotNil(t, service)

		w := httptest.NewRecorder()

		updateProfileOwnerReviewReqBody := models.SetReviewVisibilityRequest{
			Visible: boolPtr(false),
		}

		jsonPayload, err := json.Marshal(updateProfileOwnerReviewReqBody)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
		}

		updateUri := fmt.Sprintf("/api/reviews/host/set-visibility?service_id=%s", service.Data[0].ID.String())
		updateProfileOwnerReviewReq, _ := http.NewRequest("PUT", updateUri, bytes.NewBuffer(jsonPayload))
		updateProfileOwnerReviewReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		updateProfileOwnerReviewReq.Header.Set("Content-Type", "application/json")

		reviewRouter.ServeHTTP(w, updateProfileOwnerReviewReq)

		assert.Equal(t, http.StatusOK, w.Code)

		basicUser := generateUser(random, authRouter, t, "guru")
		basicUserAccessTokenCookie, _ := loginUserGetAccessToken(t, basicUser.Password, basicUser.TelegramUserID, authRouter)

		w = httptest.NewRecorder()
		getServiceReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/services/%s/service/%s", profile.Data.ID.String(), service.Data[0].ID.String()), nil)
		getServiceReq.AddCookie(&http.Cookie{Name: basicUserAccessTokenCookie.Name, Value: basicUserAccessTokenCookie.Value})
		getServiceReq.Header.Set("Content-Type", "application/json")

		serviceRouter.ServeHTTP(w, getServiceReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var servicesResponse ServiceResponse
		err = json.Unmarshal(w.Body.Bytes(), &servicesResponse)

		assert.NoError(t, err)

		assert.Equal(t, servicesResponse.Status, "success")

		assert.Equal(t, *updateProfileOwnerReviewReqBody.Visible, servicesResponse.Data.ProfileRating.ReviewTextVisible)

	})

	// -> client

	t.Run("PUT /api/reviews/client?service_id=:serviceID: profile owner can't update hist review after 48 hours passed", func(t *testing.T) {

		profileOwner := generateUser(random, authRouter, t, "")
		clientUser := generateUser(random, authRouter, t, "")

		accessTokenCookie, _ := loginUserGetAccessToken(t, profileOwner.Password, profileOwner.TelegramUserID, authRouter)

		clientAccessTokenCookie, _ := loginUserGetAccessToken(t, clientUser.Password, clientUser.TelegramUserID, authRouter)

		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, profileOwner.ID.String())

		payload := &models.CreateServiceRequest{
			ClientUserID:        clientUser.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileOwnerID:       profileOwner.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
			ProfileRating: &models.CreateProfileRatingRequest{
				Review: "I like the service! It's very good",
				Score:  ptr(5),
				RatedProfileTags: []models.CreateRatedProfileTagRequest{
					{
						Type:  "like",
						TagID: profileTags[0].ID,
					},
					{
						Type:  "like",
						TagID: profileTags[1].ID,
					},
				},
			},
			UserRating: &models.CreateUserRatingRequest{
				Review: "I liked the client! He is very kind",
				Score:  ptr(5),
				RatedUserTags: []models.CreateRatedUserTagRequest{
					{
						Type:  "like",
						TagID: userTags[0].ID,
					},
					{
						Type:  "dislike",
						TagID: userTags[1].ID,
					},
				},
			},
		}

		service, _ := createServiceFromPayload(t, *payload, serviceRouter, accessTokenCookie)

		assert.NotNil(t, service)

		newCreatedAt := time.Now().UTC().Add(-49 * time.Hour)

		result := sc.DB.Model(&models.UserRating{}).
			Where("id = ?", service.Data[0].ClientUserRatingID).
			UpdateColumn("created_at", newCreatedAt)

		if result.Error != nil {
			panic(result.Error.Error())
		}

		w := httptest.NewRecorder()

		updateClientReviewReqBody := &models.CreateUserRatingRequest{
			Review: "I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life",
			Score:  ptr(1),
			RatedUserTags: []models.CreateRatedUserTagRequest{
				{
					Type:  "dislike",
					TagID: userTags[0].ID,
				},
			},
		}

		jsonPayload, err := json.Marshal(updateClientReviewReqBody)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
		}

		updateUri := fmt.Sprintf("/api/reviews/client?service_id=%s", service.Data[0].ID.String())
		updateClientReviewReq, _ := http.NewRequest("PUT", updateUri, bytes.NewBuffer(jsonPayload))
		updateClientReviewReq.AddCookie(&http.Cookie{Name: clientAccessTokenCookie.Name, Value: clientAccessTokenCookie.Value})
		updateClientReviewReq.Header.Set("Content-Type", "application/json")

		reviewRouter.ServeHTTP(w, updateClientReviewReq)

		assert.Equal(t, http.StatusForbidden, w.Code)

	})

	t.Run("PUT /api/reviews/client?service_id=:serviceID: profile owner can update hist review before 48 hours passed", func(t *testing.T) {

		profileOwner := generateUser(random, authRouter, t, "")
		clientUser := generateUser(random, authRouter, t, "")

		accessTokenCookie, _ := loginUserGetAccessToken(t, profileOwner.Password, profileOwner.TelegramUserID, authRouter)

		clientAccessTokenCookie, _ := loginUserGetAccessToken(t, clientUser.Password, clientUser.TelegramUserID, authRouter)

		profile, _ := createProfile(t, random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors,
			intimateHairCuts, accessTokenCookie, profileRouter, profileOwner.ID.String())

		payload := &models.CreateServiceRequest{
			ClientUserID:        clientUser.ID,
			ClientUserLatitude:  floatPtr(43.259769),
			ClientUserLongitude: floatPtr(76.935246),

			ProfileID:            profile.Data.ID,
			ProfileOwnerID:       profileOwner.ID,
			ProfileUserLatitude:  floatPtr(43.259879),
			ProfileUserLongitude: floatPtr(76.934604),
			ProfileRating: &models.CreateProfileRatingRequest{
				Review: "I like the service! It's very good",
				Score:  ptr(5),
				RatedProfileTags: []models.CreateRatedProfileTagRequest{
					{
						Type:  "like",
						TagID: profileTags[0].ID,
					},
					{
						Type:  "like",
						TagID: profileTags[1].ID,
					},
				},
			},
			UserRating: &models.CreateUserRatingRequest{
				Review: "I liked the client! He is very kind",
				Score:  ptr(5),
				RatedUserTags: []models.CreateRatedUserTagRequest{
					{
						Type:  "like",
						TagID: userTags[0].ID,
					},
					{
						Type:  "dislike",
						TagID: userTags[1].ID,
					},
				},
			},
		}

		service, _ := createServiceFromPayload(t, *payload, serviceRouter, accessTokenCookie)

		assert.NotNil(t, service)

		newCreatedAt := time.Now().UTC().Add(-24 * time.Hour)

		result := sc.DB.Model(&models.UserRating{}).
			Where("id = ?", service.Data[0].ClientUserRatingID).
			UpdateColumn("created_at", newCreatedAt)

		if result.Error != nil {
			panic(result.Error.Error())
		}

		w := httptest.NewRecorder()

		updateClientReviewReqBody := &models.CreateUserRatingRequest{
			Review: "I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life",
			Score:  ptr(1),
			RatedUserTags: []models.CreateRatedUserTagRequest{
				{
					Type:  "dislike",
					TagID: userTags[0].ID,
				},
			},
		}

		jsonPayload, err := json.Marshal(updateClientReviewReqBody)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
		}

		updateUri := fmt.Sprintf("/api/reviews/client?profile_id=%s&service_id=%s", profile.Data.ID.String(), service.Data[0].ID.String())
		updateClientReviewReq, _ := http.NewRequest("PUT", updateUri, bytes.NewBuffer(jsonPayload))
		updateClientReviewReq.AddCookie(&http.Cookie{Name: clientAccessTokenCookie.Name, Value: clientAccessTokenCookie.Value})
		updateClientReviewReq.Header.Set("Content-Type", "application/json")

		reviewRouter.ServeHTTP(w, updateClientReviewReq)

		assert.Equal(t, http.StatusOK, w.Code)

		basicUser := generateUser(random, authRouter, t, "guru")
		basicUserAccessTokenCookie, _ := loginUserGetAccessToken(t, basicUser.Password, basicUser.TelegramUserID, authRouter)

		w = httptest.NewRecorder()
		getServiceReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/services/%s/service/%s", profile.Data.ID.String(), service.Data[0].ID.String()), nil)
		getServiceReq.AddCookie(&http.Cookie{Name: basicUserAccessTokenCookie.Name, Value: basicUserAccessTokenCookie.Value})
		getServiceReq.Header.Set("Content-Type", "application/json")

		serviceRouter.ServeHTTP(w, getServiceReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var servicesResponse ServiceResponse
		err = json.Unmarshal(w.Body.Bytes(), &servicesResponse)

		assert.NoError(t, err)

		assert.Equal(t, servicesResponse.Status, "success")

		assert.Equal(t, updateClientReviewReqBody.Review, servicesResponse.Data.ClientUserRating.Review)
		assert.Equal(t, updateClientReviewReqBody.Score, servicesResponse.Data.ClientUserRating.Score)

	})
}
