package main

import (
	"github.com/ivegotanidea/golang-gorm-postgres/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
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

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
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

	// todo: add profile, service controllers

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	docs.SwaggerInfo.BasePath = "/api/v1"
	router := server.Group("/api/v1")
	router.HEAD("/healthchecker", healthCheckHandler)
	router.GET("/healthchecker", healthCheckHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
