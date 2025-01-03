package main

import (
	"github.com/ivegotanidea/golang-gorm-postgres/docs"
	"github.com/ivegotanidea/golang-gorm-postgres/models"
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

	ImageController      controllers.ImageController
	ImageRouteController routes.ImageRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.InitCasbin(&config)
	initializers.ConnectDB(&config)
	initializers.Migrate()

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	ProfileController = controllers.NewProfileController(config.ParsedBaseUrl, initializers.DB)
	ProfileRouteController = routes.NewRouteProfileController(ProfileController)

	ServiceController = controllers.NewServiceController(initializers.DB, config.ReviewUpdateLimitHours)
	ServiceRouteController = routes.NewRouteServiceController(ServiceController)

	ReviewsRouteController = routes.NewRouteReviewController(ServiceController)

	DictionaryController = controllers.NewDictionaryController(initializers.DB)
	DictionaryRouteController = routes.NewRouteDictionaryController(DictionaryController)

	ImageController = controllers.NewImageController(
		initializers.DB,
		config.ImgProxyBaseUrl,
		config.S3Endpoint,
		config.ImgProxySigningHexKey,
		config.ImgProxySigningSaltHex,
		models.S3Config{
			AccessKey:    config.S3AccessKey,
			AccessSecret: config.S3AccessSecret,
			Bucket:       config.S3Bucket,
			Region:       config.S3Region,
			Endpoint:     config.S3Endpoint,
		},
		config.ProcessingGoroutinesCount)

	ImageRouteController = routes.NewRouteImageController(ImageController)

	server = gin.Default()
}

func healthCheckHandler(ctx *gin.Context) {

	if ctx.Request.Method == "GET" {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "It's alive!"})
	} else if ctx.Request.Method == "HEAD" {
		ctx.Status(http.StatusOK)
	}
}

func pingPongHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	router := server.Group("/")

	router.HEAD("/health", healthCheckHandler)
	router.GET("/health", healthCheckHandler)
	router.GET("/ping", pingPongHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := router.Group("/api/v1")

	AuthRouteController.AuthRoute(apiRouter)
	UserRouteController.UserRoute(apiRouter)
	ProfileRouteController.ProfileRoute(apiRouter)
	ServiceRouteController.ServiceRoute(apiRouter)
	ReviewsRouteController.ReviewsRoute(apiRouter)
	DictionaryRouteController.DictionaryRoute(apiRouter)
	ImageRouteController.ImageRoute(apiRouter)

	log.Fatal(server.Run(":" + config.ServerPort))
}
