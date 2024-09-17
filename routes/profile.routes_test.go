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

func SetupPCRouter(profileController *controllers.ProfileController) *gin.Engine {
	r := gin.Default()

	userRouteController := NewRouteProfileController(*profileController)

	api := r.Group("/api")
	userRouteController.ProfileRoute(api)

	return r
}

func SetupPCController() controllers.ProfileController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.InitCasbin(&config)

	profileController := controllers.NewProfileController(initializers.DB)
	profileController.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	if err := profileController.DB.AutoMigrate(
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

	return profileController
}

func TestProfileRoutes(t *testing.T) {

	ac := SetupAuthController()

	uc := SetupUCController()
	pc := SetupPCController()

	authRouter := SetupACRouter(&ac)
	userRouter := SetupUCRouter(&uc)
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

	t.Run("POST /api/profiles/: fail without access token ", func(t *testing.T) {

		w := httptest.NewRecorder()

		payload := &models.CreateProfileRequest{
			Name: "Alice",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.Header.Set("Content-Type", "application/json")
		profileRouter.ServeHTTP(w, createProfileReq)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("POST /api/profiles/: fail with access_token but bad json (only name) ", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := &models.CreateProfileRequest{
			Name: "Alice",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("POST /api/profiles/: success with access_token", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		payload.BodyTypeID = nil

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
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
			user.ID.String(), payload, profileResponse, true, false, false)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("POST /api/profiles/: fail with access_token / admin", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		_ = assignRole(initializers.DB, t, authRouter, userRouter, user.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		payload.BodyTypeID = nil

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("POST /api/profiles/: fail with access_token / moderator", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		_ = assignRole(initializers.DB, t, authRouter, userRouter, user.ID.String(), "moderator")

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		payload.BodyTypeID = nil

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("PUT /api/profiles/my/id: success self profile update", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		updatePayload := &models.UpdateOwnProfileRequest{
			Active: boolPtr(false),
			Name:   fmt.Sprintf("%s-new", payload.Name),
		}

		jsonPayload, err = json.Marshal(updatePayload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		updateProfileReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("/api/profiles/my/%s",
				profileResponse.Data.ID.String()),
			bytes.NewBuffer(jsonPayload))

		updateProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		updateProfileReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, updateProfileReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var updatedProfileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &updatedProfileResponse)

		if err != nil {
			fmt.Println("Error un marshaling response:", err)
			return
		}

		assert.Equal(t, updatedProfileResponse.Status, "success")

		getMyProfilesReq, _ := http.NewRequest("GET", "/api/profiles/my", nil)
		getMyProfilesReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		getMyProfilesReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, getMyProfilesReq)

		var profilesResponse ProfilesResponse
		err = json.Unmarshal(w.Body.Bytes(), &profilesResponse)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, 1, profilesResponse.Length)
		assert.Len(t, profilesResponse.Data, profilesResponse.Length)

		assert.Equal(t, updatePayload.Name, profilesResponse.Data[0].Name)
		assert.Equal(t, *updatePayload.Active, profilesResponse.Data[0].Active)

	})

	t.Run("PUT /api/profiles/update/id: success other user's profile update / admin", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		moderator := generateUser(random, authRouter, t)
		assignModeratorRoleResponse := assignRole(initializers.DB, t, authRouter, userRouter, moderator.ID.String(), "admin")

		assert.Equal(t, "admin", assignModeratorRoleResponse.Role)

		moderatorAccessTokenCookie, _ := loginUserGetAccessToken(t, moderator.Password, moderator.TelegramUserID, authRouter)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		updatePayload := &models.UpdateProfileRequest{
			Active:    boolPtr(false),
			Name:      fmt.Sprintf("%s-new", payload.Name),
			Moderated: boolPtr(true),
			Verified:  boolPtr(false),
		}

		jsonPayload, err = json.Marshal(updatePayload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		updateProfileReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("/api/profiles/update/%s",
				profileResponse.Data.ID.String()),
			bytes.NewBuffer(jsonPayload))

		updateProfileReq.AddCookie(&http.Cookie{Name: moderatorAccessTokenCookie.Name, Value: moderatorAccessTokenCookie.Value})
		updateProfileReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, updateProfileReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var updatedProfileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &updatedProfileResponse)

		if err != nil {
			fmt.Println("Error un marshaling response:", err)
			return
		}

		assert.Equal(t, updatedProfileResponse.Status, "success")

		getMyProfilesReq, _ := http.NewRequest("GET", "/api/profiles/my", nil)
		getMyProfilesReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		getMyProfilesReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, getMyProfilesReq)

		var profilesResponse ProfilesResponse
		err = json.Unmarshal(w.Body.Bytes(), &profilesResponse)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, 1, profilesResponse.Length)
		assert.Len(t, profilesResponse.Data, profilesResponse.Length)

		assert.Equal(t, updatePayload.Name, profilesResponse.Data[0].Name)
		assert.Equal(t, *updatePayload.Active, profilesResponse.Data[0].Active)
		assert.Equal(t, *updatePayload.Verified, profilesResponse.Data[0].Verified)
		assert.Equal(t, *updatePayload.Moderated, profilesResponse.Data[0].Moderated)
		assert.Equal(t, moderator.ID.String(), profilesResponse.Data[0].VerifiedBy.String())
		assert.Equal(t, moderator.ID.String(), profilesResponse.Data[0].ModeratedBy.String())

	})

	t.Run("PUT /api/profiles/update/id: success other user's profile update / moderator", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		moderator := generateUser(random, authRouter, t)
		assignModeratorRoleResponse := assignRole(initializers.DB, t, authRouter, userRouter, moderator.ID.String(), "moderator")

		assert.Equal(t, "moderator", assignModeratorRoleResponse.Role)

		moderatorAccessTokenCookie, _ := loginUserGetAccessToken(t, moderator.Password, moderator.TelegramUserID, authRouter)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		updatePayload := &models.UpdateProfileRequest{
			Active:    boolPtr(false),
			Name:      fmt.Sprintf("%s-new", payload.Name),
			Moderated: boolPtr(true),
			Verified:  boolPtr(false),
		}

		jsonPayload, err = json.Marshal(updatePayload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		updateProfileReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("/api/profiles/update/%s",
				profileResponse.Data.ID.String()),
			bytes.NewBuffer(jsonPayload))

		updateProfileReq.AddCookie(&http.Cookie{Name: moderatorAccessTokenCookie.Name, Value: moderatorAccessTokenCookie.Value})
		updateProfileReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, updateProfileReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var updatedProfileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &updatedProfileResponse)

		if err != nil {
			fmt.Println("Error un marshaling response:", err)
			return
		}

		assert.Equal(t, updatedProfileResponse.Status, "success")

		getMyProfilesReq, _ := http.NewRequest("GET", "/api/profiles/my", nil)
		getMyProfilesReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		getMyProfilesReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, getMyProfilesReq)

		var profilesResponse ProfilesResponse
		err = json.Unmarshal(w.Body.Bytes(), &profilesResponse)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, 1, profilesResponse.Length)
		assert.Len(t, profilesResponse.Data, profilesResponse.Length)

		assert.Equal(t, updatePayload.Name, profilesResponse.Data[0].Name)
		assert.Equal(t, *updatePayload.Active, profilesResponse.Data[0].Active)
		assert.Equal(t, *updatePayload.Verified, profilesResponse.Data[0].Verified)
		assert.Equal(t, *updatePayload.Moderated, profilesResponse.Data[0].Moderated)
		assert.Equal(t, moderator.ID.String(), profilesResponse.Data[0].VerifiedBy.String())
		assert.Equal(t, moderator.ID.String(), profilesResponse.Data[0].ModeratedBy.String())

	})

	t.Run("PUT /api/profiles/update/id: fail updating other user's profile / user", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		moderator := generateUser(random, authRouter, t)

		moderatorAccessTokenCookie, _ := loginUserGetAccessToken(t, moderator.Password, moderator.TelegramUserID, authRouter)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		updatePayload := &models.UpdateProfileRequest{
			Active:    boolPtr(false),
			Name:      fmt.Sprintf("%s-new", payload.Name),
			Moderated: boolPtr(true),
			Verified:  boolPtr(false),
		}

		jsonPayload, err = json.Marshal(updatePayload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		updateProfileReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("/api/profiles/update/%s",
				profileResponse.Data.ID.String()),
			bytes.NewBuffer(jsonPayload))

		updateProfileReq.AddCookie(&http.Cookie{Name: moderatorAccessTokenCookie.Name, Value: moderatorAccessTokenCookie.Value})
		updateProfileReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, updateProfileReq)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("GET /api/profiles: success query other user's profile / user:expert", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		tx := initializers.DB.Model(&models.User{}).Where("id = ?", secondUser.ID).Update("tier", "expert")
		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		secondUserAccessTokenCookie, _ := loginUserGetAccessToken(t, secondUser.Password, secondUser.TelegramUserID, authRouter)

		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		query := models.FindProfilesQuery{
			CityID: &payload.CityID,
		}

		queryPayload, queryErr := json.Marshal(query)
		if queryErr != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		queryProfilesReq, _ := http.NewRequest(
			"GET",
			"/api/profiles?page=1&limit=10",
			bytes.NewBuffer(queryPayload))

		queryProfilesReq.AddCookie(&http.Cookie{
			Name:  secondUserAccessTokenCookie.Name,
			Value: secondUserAccessTokenCookie.Value})

		queryProfilesReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, queryProfilesReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var profilesResponse ProfilesResponse
		err = json.Unmarshal(w.Body.Bytes(), &profilesResponse)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, 1, profilesResponse.Length)
		assert.Len(t, profilesResponse.Data, profilesResponse.Length)
	})

	t.Run("GET /api/profiles: success query other user's profile / user:guru", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		tx := initializers.DB.Model(&models.User{}).Where("id = ?", secondUser.ID).Update("tier", "guru")
		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		secondUserAccessTokenCookie, _ := loginUserGetAccessToken(t, secondUser.Password, secondUser.TelegramUserID, authRouter)

		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		query := models.FindProfilesQuery{
			CityID: &payload.CityID,
		}

		queryPayload, queryErr := json.Marshal(query)
		if queryErr != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		queryProfilesReq, _ := http.NewRequest(
			"GET",
			"/api/profiles?page=1&limit=10",
			bytes.NewBuffer(queryPayload))

		queryProfilesReq.AddCookie(&http.Cookie{
			Name:  secondUserAccessTokenCookie.Name,
			Value: secondUserAccessTokenCookie.Value})

		queryProfilesReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, queryProfilesReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var profilesResponse ProfilesResponse
		err = json.Unmarshal(w.Body.Bytes(), &profilesResponse)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, 1, profilesResponse.Length)
		assert.Len(t, profilesResponse.Data, profilesResponse.Length)
	})

	t.Run("GET /api/profiles: fail query other user's profile / user:basic", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		secondUserAccessTokenCookie, _ := loginUserGetAccessToken(t, secondUser.Password, secondUser.TelegramUserID, authRouter)

		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		query := models.FindProfilesQuery{
			CityID: &payload.CityID,
		}

		queryPayload, queryErr := json.Marshal(query)
		if queryErr != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		queryProfilesReq, _ := http.NewRequest(
			"GET",
			"/api/profiles?page=1&limit=10",
			bytes.NewBuffer(queryPayload))

		queryProfilesReq.AddCookie(&http.Cookie{
			Name:  secondUserAccessTokenCookie.Name,
			Value: secondUserAccessTokenCookie.Value})

		queryProfilesReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, queryProfilesReq)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("GET /api/profiles: fail list profiles / user: all tiers", func(t *testing.T) {

		tiers := []string{"basic", "expert", "guru"}

		for _, tier := range tiers {

			user := generateUser(random, authRouter, t)
			secondUser := generateUser(random, authRouter, t)

			tx := initializers.DB.Model(&models.User{}).Where("id = ?", secondUser.ID).Update("tier", tier)
			assert.NoError(t, tx.Error)
			assert.Equal(t, int64(1), tx.RowsAffected)

			accessTokenCookie, _ := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
			secondUserAccessTokenCookie, _ := loginUserGetAccessToken(t, secondUser.Password, secondUser.TelegramUserID, authRouter)

			for i := 0; i < 2; i++ {
				w := httptest.NewRecorder()

				payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

				jsonPayload, err := json.Marshal(payload)
				if err != nil {
					fmt.Println("Error marshaling payload:", err)
					return
				}

				createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
				createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
				createProfileReq.Header.Set("Content-Type", "application/json")

				profileRouter.ServeHTTP(w, createProfileReq)

				assert.Equal(t, http.StatusCreated, w.Code)
			}

			queryProfilesReq, _ := http.NewRequest(
				"GET",
				"/api/profiles/all?page=1&limit=10",
				nil)

			queryProfilesReq.AddCookie(&http.Cookie{
				Name:  secondUserAccessTokenCookie.Name,
				Value: secondUserAccessTokenCookie.Value})

			queryProfilesReq.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			profileRouter.ServeHTTP(w, queryProfilesReq)

			assert.Equal(t, http.StatusForbidden, w.Code)
		}
	})

	t.Run("GET /api/profiles: success list profiles / moderator", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		_ = assignRole(initializers.DB, t, authRouter, userRouter, secondUser.ID.String(), "moderator")

		accessTokenCookie, _ := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		secondUserAccessTokenCookie, _ := loginUserGetAccessToken(t, secondUser.Password, secondUser.TelegramUserID, authRouter)

		for i := 0; i < 3; i++ {
			w := httptest.NewRecorder()

			payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

			jsonPayload, err := json.Marshal(payload)
			if err != nil {
				fmt.Println("Error marshaling payload:", err)
				return
			}

			createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
			createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
			createProfileReq.Header.Set("Content-Type", "application/json")

			profileRouter.ServeHTTP(w, createProfileReq)

			assert.Equal(t, http.StatusCreated, w.Code)

			if i == 2 {

				var profileResponse CreateProfileResponse
				err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

				updatePayload := &models.UpdateOwnProfileRequest{
					Active: boolPtr(false),
					Name:   fmt.Sprintf("%s-new", payload.Name),
				}

				jsonPayload, err = json.Marshal(updatePayload)
				if err != nil {
					fmt.Println("Error marshaling payload:", err)
					return
				}

				updateProfileReq, _ := http.NewRequest(
					"PUT",
					fmt.Sprintf("/api/profiles/update/%s",
						profileResponse.Data.ID.String()),
					bytes.NewBuffer(jsonPayload))

				updateProfileReq.AddCookie(&http.Cookie{Name: secondUserAccessTokenCookie.Name, Value: secondUserAccessTokenCookie.Value})
				updateProfileReq.Header.Set("Content-Type", "application/json")

				w = httptest.NewRecorder()
				profileRouter.ServeHTTP(w, updateProfileReq)

				assert.Equal(t, http.StatusOK, w.Code)
			}
		}

		queryProfilesReq, _ := http.NewRequest(
			"GET",
			"/api/profiles/all?page=1&limit=10",
			nil)

		queryProfilesReq.AddCookie(&http.Cookie{
			Name:  secondUserAccessTokenCookie.Name,
			Value: secondUserAccessTokenCookie.Value})

		queryProfilesReq.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		profileRouter.ServeHTTP(w, queryProfilesReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var profilesResponse ProfilesResponse
		err := json.Unmarshal(w.Body.Bytes(), &profilesResponse)

		if err != nil {
			panic(err)
		}

		assert.True(t, profilesResponse.Length >= 3)

		foundInactive := false
		for _, profile := range profilesResponse.Data {
			if !profile.Active {
				foundInactive = true
				break
			}
		}

		assert.True(t, foundInactive)

	})

	t.Run("GET /api/profiles/phoneId: fail for non logged user", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)

		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		findProfileByPhoneReq, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/profiles/%s", payload.Phone),
			nil)

		findProfileByPhoneReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, findProfileByPhoneReq)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, "{\"message\":\"You are not logged in\",\"status\":\"fail\"}", w.Body.String())
	})

	t.Run("GET /api/profiles/phoneId: success for logged user", func(t *testing.T) {
		user := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)
		secondUserAccessTokenCookie, _ := loginUserGetAccessToken(t, secondUser.Password, secondUser.TelegramUserID, authRouter)

		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		findProfileByPhoneReq, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/profiles/%s", payload.Phone),
			nil)

		findProfileByPhoneReq.AddCookie(
			&http.Cookie{
				Name:  secondUserAccessTokenCookie.Name,
				Value: secondUserAccessTokenCookie.Value})

		findProfileByPhoneReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, findProfileByPhoneReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var findProfileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &findProfileResponse)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, payload.Phone, findProfileResponse.Data.Phone)
		checkProfilesMatch(t, user.ID.String(),
			payload, findProfileResponse, true, false, false)

	})

	t.Run("DELETE /api/profiles/id: success for logged user", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)

		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		findProfileByPhoneReq, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/profiles/%s", payload.Phone),
			nil)

		findProfileByPhoneReq.AddCookie(
			&http.Cookie{
				Name:  accessTokenCookie.Name,
				Value: accessTokenCookie.Value})

		findProfileByPhoneReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, findProfileByPhoneReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var findProfileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &findProfileResponse)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, payload.Phone, findProfileResponse.Data.Phone)
		checkProfilesMatch(t, user.ID.String(),
			payload, findProfileResponse, true, false, false)

		deleteProfileReq, _ := http.NewRequest(
			"DELETE",
			fmt.Sprintf("/api/profiles/%s", findProfileResponse.Data.ID.String()),
			bytes.NewBuffer(jsonPayload))

		deleteProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		deleteProfileReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, deleteProfileReq)

		assert.Equal(t, http.StatusNoContent, w.Code)
		assert.Empty(t, w.Body.String())

		getMyProfilesReq, _ := http.NewRequest("GET", "/api/profiles/my", nil)
		getMyProfilesReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		getMyProfilesReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, getMyProfilesReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var profilesResponse ProfilesResponse
		err = json.Unmarshal(w.Body.Bytes(), &profilesResponse)

		assert.Equal(t, 0, profilesResponse.Length)

	})

	t.Run("DELETE /api/profiles/id: fail for non logged user", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)

		w := httptest.NewRecorder()

		payload := generateCreateProfileRequest(random, cities, ethnos, profileTags, bodyArts, bodyTypes, hairColors, intimateHairCuts)

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		createProfileReq, _ := http.NewRequest("POST", "/api/profiles/", bytes.NewBuffer(jsonPayload))
		createProfileReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		createProfileReq.Header.Set("Content-Type", "application/json")

		profileRouter.ServeHTTP(w, createProfileReq)

		var profileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &profileResponse)

		findProfileByPhoneReq, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/api/profiles/%s", payload.Phone),
			nil)

		findProfileByPhoneReq.AddCookie(
			&http.Cookie{
				Name:  accessTokenCookie.Name,
				Value: accessTokenCookie.Value})

		findProfileByPhoneReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, findProfileByPhoneReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var findProfileResponse CreateProfileResponse
		err = json.Unmarshal(w.Body.Bytes(), &findProfileResponse)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, payload.Phone, findProfileResponse.Data.Phone)
		checkProfilesMatch(t, user.ID.String(),
			payload, findProfileResponse, true, false, false)

		deleteProfileReq, _ := http.NewRequest(
			"DELETE",
			fmt.Sprintf("/api/profiles/%s", findProfileResponse.Data.ID.String()),
			bytes.NewBuffer(jsonPayload))

		deleteProfileReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		profileRouter.ServeHTTP(w, deleteProfileReq)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, "{\"message\":\"You are not logged in\",\"status\":\"fail\"}", w.Body.String())
	})
}
