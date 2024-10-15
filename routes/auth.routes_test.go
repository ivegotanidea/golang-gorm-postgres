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

// SetupACRouter sets up the router for testing.
func SetupACRouter(authController *controllers.AuthController) *gin.Engine {
	r := gin.Default()

	authRouteController := NewAuthRouteController(*authController)

	api := r.Group("/api")
	authRouteController.AuthRoute(api)

	return r
}

func SetupAuthController() controllers.AuthController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	authController := controllers.NewAuthController(initializers.DB)
	authController.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Migrate the schema
	if err := authController.DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Service{}, &models.Photo{}, &models.ProfileOption{}, &models.UserRating{}, &models.ProfileRating{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return authController
}

func TestAuthRoutes(t *testing.T) {

	ac := SetupAuthController()
	router := SetupACRouter(&ac)
	random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))

	t.Cleanup(func() {
		utils.CleanupTestUsers(ac.DB)
		//utils.DropAllTables(ac.DB)
	})

	t.Run("POST /api/auth/register: successful registration OK ", func(t *testing.T) {
		name := "testiculous-andrew"
		phone := utils.GenerateRandomPhoneNumber(random, 0)
		telegramUserId := fmt.Sprintf("%d", rand.Int64())

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.NoError(t, err)

		user := jsonResponse["data"].(map[string]interface{})

		assert.Equal(t, "success", jsonResponse["status"])
		// Check name and phone
		assert.Equal(t, name, user["name"])
		assert.Equal(t, phone, user["phone"])
	})

	t.Run("POST /api/auth/register: empty name registration FAIL ", func(t *testing.T) {
		name := ""
		phone := utils.GenerateRandomPhoneNumber(random, 0)

		telegramUserId := fmt.Sprintf("%d", rand.Int64())
		errMessage := "Key: 'BotSignUpInput.Name' Error:Field validation for 'Name' failed on the 'required' tag"

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		message := jsonResponse["message"]
		assert.NoError(t, err)
		assert.Equal(t, "error", jsonResponse["status"])

		// Check name and phone
		assert.Equal(t, errMessage, message)
	})

	t.Run("POST /api/auth/register: empty phone registration FAIL ", func(t *testing.T) {
		name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
		phone := ""

		telegramUserId := fmt.Sprintf("%d", rand.Int64())
		errMessage := "Key: 'BotSignUpInput.Phone' Error:Field validation for 'Phone' failed on the 'required' tag"

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		message := jsonResponse["message"]
		assert.NoError(t, err)
		assert.Equal(t, "error", jsonResponse["status"])

		// Check name and phone
		assert.Equal(t, errMessage, message)
	})

	t.Run("POST /api/auth/register: phone of 10 symbols registration FAIL ", func(t *testing.T) {
		name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
		phone := utils.GenerateRandomPhoneNumber(random, 10)

		telegramUserId := fmt.Sprintf("%d", rand.Int64())
		errMessage := "Key: 'BotSignUpInput.Phone' Error:Field validation for 'Phone' failed on the 'min' tag"

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		message := jsonResponse["message"]
		assert.NoError(t, err)
		assert.Equal(t, "error", jsonResponse["status"])

		// Check name and phone
		assert.Equal(t, errMessage, message)
	})

	t.Run("POST /api/auth/register: phone of 12 symbols registration FAIL ", func(t *testing.T) {
		name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
		phone := utils.GenerateRandomPhoneNumber(random, 12)

		telegramUserId := fmt.Sprintf("%d", rand.Int64())
		errMessage := "Key: 'BotSignUpInput.Phone' Error:Field validation for 'Phone' failed on the 'max' tag"

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		message := jsonResponse["message"]
		assert.NoError(t, err)
		assert.Equal(t, "error", jsonResponse["status"])

		// Check name and phone
		assert.Equal(t, errMessage, message)
	})

	t.Run("POST /api/auth/register: empty telegramId registration FAIL ", func(t *testing.T) {
		name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
		phone := utils.GenerateRandomPhoneNumber(random, 0)
		telegramUserId := ""
		errMessage := "Key: 'BotSignUpInput.TelegramUserId' Error:Field validation for 'TelegramUserId' failed on the 'required' tag"

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		message := jsonResponse["message"]
		assert.NoError(t, err)
		assert.Equal(t, "error", jsonResponse["status"])

		// Check name and phone
		assert.Equal(t, errMessage, message)
	})

	t.Run("POST /api/auth/register: empty registration data FAIL ", func(t *testing.T) {
		name := ""
		phone := ""
		telegramUserId := ""
		errMessage := "Key: 'BotSignUpInput.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'BotSignUpInput.Phone' Error:Field validation for 'Phone' failed on the 'required' tag\nKey: 'BotSignUpInput.TelegramUserId' Error:Field validation for 'TelegramUserId' failed on the 'required' tag"

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		message := jsonResponse["message"]
		assert.NoError(t, err)
		assert.Equal(t, "error", jsonResponse["status"])

		// Check name and phone
		assert.Equal(t, errMessage, message)
	})

	t.Run("POST /api/auth/login + GET api/auth/refresh + GET api/auth/logout", func(t *testing.T) {
		name := utils.GenerateRandomStringWithPrefix(random, 10, "test-")
		phone := utils.GenerateRandomPhoneNumber(random, 0)
		telegramUserId := fmt.Sprintf("%d", rand.Int64())

		payload := fmt.Sprintf(`{"name": "%s", "phone": "%s", "telegramUserId": "%s"}`, name, phone, telegramUserId)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var jsonResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		assert.NoError(t, err)

		user := jsonResponse["data"].(map[string]interface{})

		assert.Equal(t, "success", jsonResponse["status"])
		// Check name and phone
		assert.Equal(t, name, user["name"])
		assert.Equal(t, phone, user["phone"])
		assert.NotEmptyf(t, user["password"], "")

		w = httptest.NewRecorder()
		password := user["password"].(string)
		payloadLogin := fmt.Sprintf(`{"telegramUserId": "%s", "password": "%s"}`, telegramUserId, password)
		loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer([]byte(payloadLogin)))
		loginReq.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, loginReq)

		err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
		status := jsonResponse["status"]

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, status, "success")
		assert.NotEmpty(t, jsonResponse["access_token"])

		// Extract refresh_token from cookies
		cookies := w.Result().Cookies()
		var refreshTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "refresh_token" {
				refreshTokenCookie = cookie
				break
			}
		}
		assert.NotNil(t, refreshTokenCookie)
		assert.NotEmpty(t, refreshTokenCookie.Value)

		w = httptest.NewRecorder()
		refreshReq, _ := http.NewRequest("GET", "/api/auth/refresh", nil)
		refreshReq.AddCookie(&http.Cookie{Name: refreshTokenCookie.Name, Value: refreshTokenCookie.Value})
		router.ServeHTTP(w, refreshReq)

		assert.Equal(t, http.StatusOK, w.Code)
		cookies = w.Result().Cookies()
		var accessTokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "access_token" {
				accessTokenCookie = cookie
				break
			}
		}

		assert.NotNil(t, accessTokenCookie)
		assert.NotEmpty(t, accessTokenCookie.Value)

		logoutReq, err := http.NewRequest("GET", "/api/auth/logout", nil)
		logoutReq.AddCookie(accessTokenCookie)
		logoutReq.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		router.ServeHTTP(w, logoutReq)

		assert.Equal(t, http.StatusOK, w.Code)

		reqWithClearedCookie, _ := http.NewRequest("GET", "/api/auth/refresh", nil)

		// Ensure cookies are cleared after logout
		cookies = w.Result().Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "access_token" || cookie.Name == "refresh_token" || cookie.Name == "logged_in" {
				assert.Empty(t, cookie.Value)
				assert.Equal(t, -1, cookie.MaxAge)

				reqWithClearedCookie.AddCookie(cookie)
			}
		}

		// Attempt to access a protected route after logout
		reqWithClearedCookie.Header.Set("Content-Type", "application/json")

		w = httptest.NewRecorder()
		router.ServeHTTP(w, reqWithClearedCookie)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}
