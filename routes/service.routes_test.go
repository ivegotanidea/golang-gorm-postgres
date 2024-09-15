package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestServiceRoutes(t *testing.T) {

	//ac := SetupAuthController()
	uc := SetupUCController()

	//authRouter := SetupACRouter(&ac)
	userRouter := SetupUCRouter(&uc)

	//random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))

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
}
