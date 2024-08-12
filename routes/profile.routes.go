package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/middleware"
)

type ProfileRouteController struct {
	profileController controllers.ProfileController
}

func NewRouteProfileController(profileController controllers.ProfileController) ProfileRouteController {
	return ProfileRouteController{profileController}
}

func (pc *ProfileRouteController) ProfileRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.Use(middleware.DeserializeUser())

	router.POST("/", pc.profileController.CreateProfile)
	router.GET("/", pc.profileController.FindProfiles)
	router.PUT("/:id", pc.profileController.UpdateProfile)
	router.GET("/:phone", pc.profileController.FindProfileByPhone)
	router.DELETE("/:id", pc.profileController.DeleteProfile)
}
