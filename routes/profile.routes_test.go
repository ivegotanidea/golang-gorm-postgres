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
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
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

func PopulateProfileTags(db gorm.DB) []models.ProfileTag {
	var profileTags = []string{
		"Классика",
		"Минет c/без резинки",
		"Глубокий минет с резинкой",
		"Глубокий минет c/без резинки c окончанием",
		"Разрешу куннилингус",
		"Минет c/без резинки c окончанием",
		"Минет с резинкой",
		"Массаж любительский",
		"Массаж профессиональный",
		"Вагинальный фистинг",
		"Расслабляющий массаж",
		"Поцелуи в губы",
		"Массаж простаты",
		"Классический массаж",
		"Поеду отдыхать (в клуб, ресторан и.т.д.). Вечер:",
		"Тайский боди массаж",
		"Глубокий минет c/без резинки",
		"Анилингус, побалую язычком очко",
		"Услуги Госпоже",
		"Услуги семейным парам",
		"Французский поцелуй",
		"Эротический массаж",
		"Секс по телефону",
		"Обслуживаю мальчишники. Вечер:",
		"Групповой секс",
		"Стриптиз любительский",
		"Ветка сакуры",
		"Снимусь на видео",
		"Анальный секс",
		"Ролевые игры, наряды",
		"Фото на память",
		"Сделаю минет",
		"Стриптиз профессиональный",
		"Глубокий минет без резинки c окончанием",
		"Минет без резинки c окончанием",
		"Глубокий минет без резинки",
		"Минет без резинки",
		"Обслуживаю девушек",
		"Обслуживаю парней",
		"Сделаю куннилингус",
		"Обслуживаю девишники/вечеринки. Вечер:",
		"Обслуживаю вечеринки. Вечер:",
	}

	var tags []models.ProfileTag
	for _, profileTag := range profileTags {
		tag := models.ProfileTag{
			Name: profileTag,
		}
		tags = append(tags, tag)
	}

	tx := db.Begin()
	// Batch insert options
	if err := tx.Create(&tags).Error; err != nil {
		tx.Rollback()
		panic("failed to create profile tags: " + err.Error())
	}

	return tags
}

func TestProfileRoutes(t *testing.T) {

	//ac := SetupAuthController()
	//uc := SetupUCController()
	pc := SetupPCController()

	//authRouter := SetupACRouter(&ac)
	//userRouter := SetupUCRouter(&uc)
	profileRouter := SetupPCRouter(&pc)
	//profileTags := PopulateProfileTags(*ac.DB)
	//createOwnerUser(profileController.DB)

	//random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))

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

}
