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

	t.Run("POST /api/auth/login + GET api/auth/refresh + GET api/auth/logout", func(t *testing.T) {
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
		assert.NotEmptyf(t, user["password"], "")

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
		authRouter.ServeHTTP(w, refreshReq)

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
		authRouter.ServeHTTP(w, logoutReq)

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
		authRouter.ServeHTTP(w, reqWithClearedCookie)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}
