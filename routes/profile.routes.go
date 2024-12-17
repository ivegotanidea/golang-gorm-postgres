package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/controllers"
	"github.com/ivegotanidea/golang-gorm-postgres/middleware"
)

type ProfileRouteController struct {
	profileController controllers.ProfileController
}

func NewRouteProfileController(profileController controllers.ProfileController) ProfileRouteController {
	return ProfileRouteController{profileController}
}

// @BasePath /api/v1/profile

func (pc *ProfileRouteController) ProfileRoute(rg *gin.RouterGroup) {

	router := rg.Group("profiles")

	router.POST("/", middleware.DeserializeUser(), pc.profileController.CreateProfile)

	router.GET("/my", middleware.DeserializeUser(), pc.profileController.GetMyProfiles)

	router.GET("", middleware.DeserializeUser(), middleware.AbacMiddleware("profiles", "query"), pc.profileController.FindProfiles)

	router.GET("/list", pc.profileController.ListProfilesNonAuth)
	router.GET("/all", middleware.DeserializeUser(), middleware.AbacMiddleware("profiles", "list"), pc.profileController.ListProfiles)

	router.PUT("/my/:id", middleware.DeserializeUser(), pc.profileController.UpdateOwnProfile)
	router.POST("/:id/photos", middleware.DeserializeUser(), pc.profileController.UpdateProfilePhotos)

	router.PUT("/update/:id", middleware.DeserializeUser(), middleware.AbacMiddleware("profiles", "update"), pc.profileController.UpdateProfile)

	// todo: should have captcha set
	// todo: should have rate limiter set
	router.GET("/:id", middleware.DeserializeUser(), pc.profileController.FindProfileByID)

	router.DELETE("/:id", middleware.DeserializeUser(), pc.profileController.DeleteProfile)
}
