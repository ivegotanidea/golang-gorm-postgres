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

type UserResponse struct {
	Status string `json:"status"`
	Data   struct {
		User models.UserResponse `json:"user"`
	} `json:"data"`
}

type FindUserResponse struct {
	Status string              `json:"status"`
	Data   models.UserResponse `json:"data"`
}

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

	// Migrate the schema
	if err := userController.DB.AutoMigrate(&models.User{}, &models.Profile{}); err != nil {
		panic("failed to migrate database")
	}

	return userController
}

func generateUser(random *rand.Rand, authRouter *gin.Engine, t *testing.T) models.UserResponse {
	name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
	phone := utils.GenerateRandomPhoneNumber(random, 0)
	telegramUserId := fmt.Sprintf("%d", rand.Int64())

	payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	authRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var userResponse UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &userResponse)
	assert.NoError(t, err)

	user := userResponse.Data.User

	return user
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

		data := jsonResponse["data"].(map[string]interface{})
		user := data["user"].(map[string]interface{})

		assert.Equal(t, "success", jsonResponse["status"])
		// Check name and phone
		assert.Equal(t, name, user["name"])
		assert.Equal(t, phone, user["phone"])

		w = httptest.NewRecorder()
		password := user["password"].(string)
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%s", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq)

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		w = httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/me", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		jsonResponse = make(map[string]interface{})
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, w.Code)

		data = jsonResponse["data"].(map[string]interface{})
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

		data := jsonResponse["data"].(map[string]interface{})
		user := data["user"].(map[string]interface{})

		assert.Equal(t, "success", jsonResponse["status"])
		// Check name and phone
		assert.Equal(t, name, user["name"])
		assert.Equal(t, phone, user["phone"])

		w = httptest.NewRecorder()
		password := user["password"].(string)
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%s", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq)

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		w = httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		jsonResponse = make(map[string]interface{})
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusForbidden, w.Code)

	})

	t.Run("GET /api/user: moderator, success list users", func(t *testing.T) {
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

		data := jsonResponse["data"].(map[string]interface{})
		user := data["user"].(map[string]interface{})

		assert.Equal(t, "success", jsonResponse["status"])
		// Check name and phone
		assert.Equal(t, name, user["name"])
		assert.Equal(t, phone, user["phone"])

		tx := initializers.DB.Model(&models.User{}).Where("id = ?", user["id"]).Update("tier", "moderator")

		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		w = httptest.NewRecorder()
		password := user["password"].(string)
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%s", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq)

		jsonResponse = make(map[string]interface{})
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		w = httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		jsonResponse = make(map[string]interface{})
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("GET /api/user: admin, success list users", func(t *testing.T) {
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

		data := jsonResponse["data"].(map[string]interface{})
		user := data["user"].(map[string]interface{})

		assert.Equal(t, "success", jsonResponse["status"])
		// Check name and phone
		assert.Equal(t, name, user["name"])
		assert.Equal(t, phone, user["phone"])

		tx := initializers.DB.Model(&models.User{}).Where("id = ?", user["id"]).Update("tier", "admin")

		assert.NoError(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)

		w = httptest.NewRecorder()
		password := user["password"].(string)
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%s", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq)

		jsonResponse = make(map[string]interface{})
		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		w = httptest.NewRecorder()
		meReq, _ := http.NewRequest("GET", "/api/users/", nil)
		meReq.Header.Set("Content-Type", "application/json")
		meReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, meReq)

		jsonResponse = make(map[string]interface{})
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

		data := jsonResponse["data"].(map[string]interface{})
		user := data["user"].(map[string]interface{})

		assert.NotEmpty(t, data)
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

		var jsonResponse map[string]interface{}

		w := httptest.NewRecorder()
		password := firstUser.Password
		telegramUserId := firstUser.TelegramUserID
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%d", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq)

		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.NoError(t, err)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		w = httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?phone=%s", secondUser.Phone)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse FindUserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.NotEmpty(t, userResponse)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /api/users/user: success by id with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		var jsonResponse map[string]interface{}

		w := httptest.NewRecorder()
		password := firstUser.Password
		telegramUserId := firstUser.TelegramUserID
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%d", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq)

		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.NoError(t, err)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		w = httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?id=%s", secondUser.ID)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse FindUserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.NotEmpty(t, userResponse)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /api/users/user: success by telegramId with access token", func(t *testing.T) {
		firstUser := generateUser(random, authRouter, t)
		secondUser := generateUser(random, authRouter, t)

		var jsonResponse map[string]interface{}

		w := httptest.NewRecorder()
		password := firstUser.Password
		telegramUserId := firstUser.TelegramUserID
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%d", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		authRouter.ServeHTTP(w, loginReq)

		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)

		assert.NoError(t, err)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		w = httptest.NewRecorder()

		url := fmt.Sprintf("/api/users/user?telegramUserId=%d", secondUser.TelegramUserID)
		findUserReq, _ := http.NewRequest("GET", url, nil)
		findUserReq.Header.Set("Content-Type", "application/json")
		findUserReq.AddCookie(&http.Cookie{Name: accessTokenCookie.Name, Value: accessTokenCookie.Value})
		userRouter.ServeHTTP(w, findUserReq)

		var userResponse FindUserResponse
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.Nil(t, err)
		assert.NotEmpty(t, userResponse)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
