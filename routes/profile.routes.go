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

	router := rg.Group("profiles")
	router.Use(middleware.DeserializeUser())

	router.POST("/", pc.profileController.CreateProfile)

	router.GET("/my", pc.profileController.GetMyProfiles)

	router.GET("", middleware.AbacMiddleware("profiles", "query"), pc.profileController.FindProfiles)

	router.GET("/all", pc.profileController.ListProfiles)

	router.PUT("/my/:id", pc.profileController.UpdateOwnProfile)
	router.PUT("/update/:id", middleware.AbacMiddleware("profiles", "update"), pc.profileController.UpdateProfile)

	router.GET("/:phone", pc.profileController.FindProfileByPhone)

	router.DELETE("/:id", pc.profileController.DeleteProfile)
}
