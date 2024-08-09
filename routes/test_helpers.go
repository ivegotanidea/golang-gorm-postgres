package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"gorm.io/gorm"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserResponse struct {
	Status string `json:"status"`
	Data   struct {
		User models.UserResponse `json:"user"`
	} `json:"data"`
}

func getOwnerUser() models.User {
	return models.User{
		ID:             uuid.Max,
		Name:           "He Who Remains",
		Phone:          "77778889900",
		TelegramUserId: 6794234746,
		Password:       "h5sh3d", // Ensure this is hashed
		Avatar:         "https://akm-img-a-in.tosshub.com/indiatoday/images/story/202311/tom-hiddleston-in-a-still-from-loki-2-27480244-16x9_0.jpg",
		Verified:       true,
		HasProfile:     false,
		Tier:           "owner",
	}
}

func createOwnerUser(db *gorm.DB) {

	owner := getOwnerUser()

	if err := db.Where("tier = ?", "owner").FirstOrCreate(&owner).Error; err != nil {
		panic(err)
	}
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

func loginUserGetAccessToken(t *testing.T, password string, telegramUserId int64, authRouter *gin.Engine) (*http.Cookie, error) {
	var jsonResponse map[string]interface{}

	w := httptest.NewRecorder()
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

	for _, cookie := range cookies {
		if cookie.Name == "access_token" {
			return cookie, err
		}
	}
	return nil, errors.New("cookie not found")
}
