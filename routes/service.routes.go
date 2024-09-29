package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/middleware"
)

type ServiceRouteController struct {
	serviceController controllers.ServiceController
}

func NewRouteServiceController(serviceController controllers.ServiceController) ServiceRouteController {
	return ServiceRouteController{serviceController}
}

func (sc *ServiceRouteController) ServiceRoute(rg *gin.RouterGroup) {
	router := rg.Group("services")

	router.Use(middleware.DeserializeUser())

	router.POST("/", sc.serviceController.CreateService)

	router.GET("/:profileID", sc.serviceController.GetProfileServices)

	router.GET("/all", middleware.AbacMiddleware("services", "list"), sc.serviceController.ListServices)

	reviewsRouter := router.Group("reviews")

	reviewsRouter.PUT("/client/:profileID", sc.serviceController.UpdateClientUserReviewOnProfile)
	reviewsRouter.PUT("/client/:profileID/:serviceID", sc.serviceController.HideProfileOwnerReview)

	reviewsRouter.PUT("/host/:profileID", sc.serviceController.UpdateProfileOwnerReviewOnClientUser)
	reviewsRouter.PUT("/host/:profileID/:serviceID", sc.serviceController.HideUserReview)

	// if coordinates not provided or not close enough
	// then automark as unverified
}
