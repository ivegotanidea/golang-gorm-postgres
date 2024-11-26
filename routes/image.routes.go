package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/controllers"
	"github.com/ivegotanidea/golang-gorm-postgres/middleware"
)

type ImageRouteController struct {
	imageController controllers.ImageController
}

func NewImageRouteController(imageController controllers.ImageController) ImageRouteController {
	return ImageRouteController{imageController}
}

// @BasePath /api/v1/auth

func (ic *ImageRouteController) ImageRoute(rg *gin.RouterGroup) {
	router := rg.Group("images")

	router.Use(middleware.DeserializeUser())

	router.POST("/", ic.imageController.UploadProfileImages)
}
