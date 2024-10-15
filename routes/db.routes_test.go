package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func SetupDB() controllers.AuthController {
	var err error
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	controller := controllers.NewAuthController(initializers.DB)
	controller.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Migrate the schema
	if err := controller.DB.AutoMigrate(
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

	return controller
}

func TestDbMigrations(t *testing.T) {

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
}
