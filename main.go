package main

import (
	"github.com/ivegotanidea/golang-gorm-postgres/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/controllers"
	"github.com/ivegotanidea/golang-gorm-postgres/initializers"
	"github.com/ivegotanidea/golang-gorm-postgres/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	ProfileController      controllers.ProfileController
	ProfileRouteController routes.ProfileRouteController

	ServiceController      controllers.ServiceController
	ServiceRouteController routes.ServiceRouteController

	ReviewsRouteController routes.ReviewsRouteController

	DictionaryController      controllers.DictionaryController
	DictionaryRouteController routes.DictionaryRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	ProfileController = controllers.NewProfileController(initializers.DB)
	ProfileRouteController = routes.NewRouteProfileController(ProfileController)

	ServiceController = controllers.NewServiceController(initializers.DB, config.ReviewUpdateLimitHours)
	ServiceRouteController = routes.NewRouteServiceController(ServiceController)

	ReviewsRouteController = routes.NewRouteReviewController(ServiceController)

	DictionaryController = controllers.NewDictionaryController(initializers.DB)
	DictionaryRouteController = routes.NewRouteDictionaryController(DictionaryController)

	server = gin.Default()
}

func healthCheckHandler(ctx *gin.Context) {

	var message string

	if ctx.Request.Method == "GET" {
		message = "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	} else if ctx.Request.Method == "HEAD" {
		ctx.Status(http.StatusOK)
	}
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	router := server.Group("/api/v1")

	router.HEAD("/healthchecker", healthCheckHandler)
	router.GET("/healthchecker", healthCheckHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	ProfileRouteController.ProfileRoute(router)
	ServiceRouteController.ServiceRoute(router)
	ReviewsRouteController.ReviewsRoute(router)
	DictionaryRouteController.DictionaryRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
