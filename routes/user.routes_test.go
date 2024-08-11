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

// SetupUCRouter sets up the router for testing.
func SetupUCRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	userRouteController := NewRouteUserController(*userController)

	api := r.Group("/api")
	userRouteController.UserRoute(api)

	return r
}

func SetupUCController() controllers.UserController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.InitCasbin(&config)

	userController := controllers.NewUserController(initializers.DB)
	userController.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	createOwnerUser(userController.DB)

	// Migrate the schema
	if err := userController.DB.AutoMigrate(&models.User{}, &models.Profile{}); err != nil {
		panic("failed to migrate database")
	}

	return userController
}

func TestUserRoutes(t *testing.T) {

	ac := SetupAuthController()
	uc := SetupUCController()

	authRouter := SetupACRouter(&ac)
	userRouter := SetupUCRouter(&uc)

	random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))

	t.Cleanup(func() {
		utils.CleanupTestUsers(uc.DB)
		utils.DropAllTables(uc.DB)
	})

	t.Run("GET /api/user/me: fail without access token ", func(t *testing.T) {

		w := httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/me", nil)
		meReq.Header.Set("Content-Type", "application/json")
		userRouter.ServeHTTP(w, meReq)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

	})

	t.Run("GET /api/user/me: success with access token", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/me", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, w.Code)

		data := jsonResponse["data"].(map[string]interface{})
		assert.NotEmpty(t, data)

	})

	t.Run("GET /api/user: no access_token, forbidden to list users", func(t *testing.T) {

		w := httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		userRouter.ServeHTTP(w, meReq)

		jsonResponse := make(map[string]interface{})
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

	})

	t.Run("GET /api/user: basic user, forbidden to list users", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusForbidden, w.Code)

	})

	t.Run("GET /api/user: moderator, success list users", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		user = assignRole(initializers.DB, t, authRouter, userRouter, user.ID.String(), "moderator")

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("GET /api/user: admin, success list users", func(t *testing.T) {
		user := generateUser(random, authRouter, t)

		user = assignRole(initializers.DB, t, authRouter, userRouter, user.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, user.Password, user.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("GET /api/users/user: fail without access token", func(t *testing.T) {
		name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
		phone := utils.GenerateRandomPhoneNumber(random, 0)
		telegramUserId := fmt.Sprintf("%d", rand.Int64())

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.NoError(t, err)

		user := jsonResponse["data"].(map[string]interface{})

		assert.NotEmpty(t, user)

		w = httptest.NewRecorder()

		jsonResponse = make(map[string]interface{})
		url := fmt.Sprintf("/api/users/user?phone=%s", user["phone"])
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		userRouter.ServeHTTP(w, findUserReq)

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

	})

	t.Run("GET /api/users/user: success by phone with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?phone=%s", secondUser.Phone)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse UserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.NotEmpty(t, userResponse)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /api/users/user: success by id with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?id=%s", secondUser.ID)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse UserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.NotEmpty(t, userResponse)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /api/users/user: success by telegramId with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?telegramUserId=%d", secondUser.TelegramUserID)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse UserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.NotEmpty(t, userResponse)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /api/users/user: 404 non existing phone with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?phone=%s", "7000000000")
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("GET /api/users/user: 404 non existing id with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?id=%s", "00000000-0000-0000-0000-000000000000")
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("GET /api/users/user: 404 non existing telegramId with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?telegramUserId=%d", 1991)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("DELETE /api/users/user: success with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user")
		findUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		assert.Equal(t, http.StatusNoContent, w.Code)

		var jsonResponse map[string]interface{}

		w = httptest.NewRecorder()
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%d", "password": "%s"}`, firstUser.TelegramUserID, firstUser.Password)
		loginReq2, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq2.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq2)

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, jsonResponse["status"], "fail")
		assert.Nil(t, jsonResponse["access_token"])
	})

	t.Run("DELETE /api/users/user: fail without access token", func(t *testing.T) {
		var jsonResponse map[string]interface{}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user")
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		userRouter.ServeHTTP(w, delUserReq)

		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, jsonResponse["message"], "You are not logged in")
		assert.Equal(t, jsonResponse["status"], "fail")
	})

	t.Run("DELETE /api/users/user: fail moderator deletes user", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		tx := initializers.DB.Model(&models.User{}).Where("id = ?", firstUser.ID).Update("tier", "moderator")
		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		assert.NotEmpty(t, w.Body.String())

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("DELETE /api/users/user: success admin deletes user", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		assert.Empty(t, w.Body.String())

		assert.Equal(t, http.StatusNoContent, w.Code)

		w = httptest.NewRecorder()

		findUserUrl := fmt.Sprintf("/api/users/user?id=%s", secondUser.ID)
		findUserReq, _ := http.NewRequest("GET", findUserUrl, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse UserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.Equal(t, userResponse.Status, "fail")

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("DELETE /api/users/user: fail moderator deletes moderator", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "moderator")
		secondUser = assignRole(initializers.DB, t, authRouter, userRouter, secondUser.ID.String(), "moderator")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		var jsonResponse map[string]interface{}

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("DELETE /api/users/user: fail moderator deletes admin", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		tx := initializers.DB.Model(&models.User{}).Where("id = ?", firstUser.ID).Update("tier", "moderator")
		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		tx = initializers.DB.Model(&models.User{}).Where("id = ?", secondUser.ID).Update("tier", "admin")
		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		var jsonResponse map[string]interface{}

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("DELETE /api/users/user: success admin deletes moderator", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "admin")
		secondUser = assignRole(initializers.DB, t, authRouter, userRouter, secondUser.ID.String(), "moderator")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		var jsonResponse map[string]interface{}

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, jsonResponse)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("DELETE /api/users/user: fail admin deletes admin", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "admin")
		secondUser = assignRole(initializers.DB, t, authRouter, userRouter, secondUser.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		var jsonResponse map[string]interface{}

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Equal(t, jsonResponse["status"], "fail")

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("DELETE /api/users/user: success owner deletes admin", func(t *testing.T) {
		owner := getOwnerUser()
		secondUser := generateUser(random, authRouter, t)

		secondUser = assignRole(initializers.DB, t, authRouter, userRouter, secondUser.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, owner.Password, owner.TelegramUserId, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		var jsonResponse map[string]interface{}

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, jsonResponse)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("DELETE /api/users/user: fail moderator deletes owner", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := getOwnerUser()

		tx := initializers.DB.Model(&models.User{}).Where("id = ?", firstUser.ID).Update("tier", "moderator")
		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		var jsonResponse map[string]interface{}

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("DELETE /api/users/user: fail admin deletes owner", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := getOwnerUser()

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		delUserReq, _ := http.NewRequest(http.MethodDelete, url, nil)
		delUserReq.Header.Set("Content-Type", "application/json")
		delUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, delUserReq)

		var jsonResponse map[string]interface{}

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Equal(t, jsonResponse["status"], "fail")

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("UPDATE /api/users/user: fail with access token and empty update", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user")
		updUserReq, _ := http.NewRequest(http.MethodPut, url, nil)
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusBadGateway, w.Code)

	})

	t.Run("UPDATE /api/users/user: fail without access token", func(t *testing.T) {

		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user")
		updUserReq, _ := http.NewRequest(http.MethodPut, url, nil)
		updUserReq.Header.Set("Content-Type", "application/json")
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("UPDATE /api/users/user: success with access token, update name", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUser{
			Name: fmt.Sprintf("%s-new", firstUser.Name),
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user")
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var userResponse UserResponse

		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.Equal(t, userResponse.Status, "success")
		assert.Equal(t, userResponse.Data.Name, payload.Name)
	})

	t.Run("UPDATE /api/users/user: success with access token, update phone", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUser{
			Phone: "77000000101",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user")
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var userResponse UserResponse

		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.Equal(t, userResponse.Status, "success")
		assert.Equal(t, userResponse.Data.Phone, payload.Phone)
	})

	t.Run("UPDATE /api/users/user: success with access token, update avatar", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUser{
			Avatar: "https://jollycontrarian.com/images/6/6c/Rickroll.jpg",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user")
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusOK, w.Code)

		var userResponse UserResponse

		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.Equal(t, userResponse.Status, "success")
		assert.Equal(t, userResponse.Data.Avatar, payload.Avatar)
	})

	t.Run("UPDATE /api/users/user: basic user fails with access token, update avatar", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUserPrivileged{
			Avatar: "https://jollycontrarian.com/images/6/6c/Rickroll.jpg",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("UPDATE /api/users/user: moderator success with access token, update avatar", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "moderator")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUserPrivileged{
			Avatar: "https://jollycontrarian.com/images/6/6c/Rickroll.jpg",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("UPDATE /api/users/user: admin success with access token, deactivate user", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUserPrivileged{
			Active: false,
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusOK, w.Code)

		w = httptest.NewRecorder()
		url = fmt.Sprintf("/api/users/user?phone=%s", secondUser.Phone)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse UserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.Equal(t, userResponse.Status, "success")
		assert.NotEmpty(t, userResponse)
		assert.Equal(t, userResponse.Data.Active, false)
	})

	t.Run("UPDATE /api/users/user: admin success with access token, verify user", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUserPrivileged{
			Verified: true,
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusOK, w.Code)

		w = httptest.NewRecorder()
		url = fmt.Sprintf("/api/users/user?phone=%s", secondUser.Phone)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse UserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.Equal(t, userResponse.Status, "success")
		assert.NotEmpty(t, userResponse)
		assert.Equal(t, userResponse.Data.Verified, true)
	})

	t.Run("GET /api/users/users: guru success list users with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		firstUser = assignRole(initializers.DB, t, authRouter, userRouter, firstUser.ID.String(), "admin")

		accessTokenCookie, err := loginUserGetAccessToken(t, firstUser.Password, firstUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w := httptest.NewRecorder()

		payload := &models.UpdateUserPrivileged{
			Tier: "expert",
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		url := fmt.Sprintf("/api/users/user/%s", secondUser.ID)
		updUserReq, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
		updUserReq.Header.Set("Content-Type", "application/json")
		updUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, updUserReq)

		assert.Equal(t, http.StatusOK, w.Code)

		w = httptest.NewRecorder()
		url = fmt.Sprintf("/api/users/user?phone=%s", secondUser.Phone)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse UserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.Equal(t, userResponse.Status, "success")
		assert.NotEmpty(t, userResponse)
		assert.Equal(t, "expert", userResponse.Data.Tier)

		accessTokenCookie, err = loginUserGetAccessToken(t, secondUser.Password, secondUser.TelegramUserID, authRouter)

		if err != nil {
			panic(err)
		}

		w = httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.NotEmpty(t, jsonResponse["data"])
		assert.Equal(t, http.StatusOK, w.Code)

	})
}
