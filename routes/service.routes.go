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

	router.GET("/:profileID/service/:serviceID", sc.serviceController.GetService)

	router.GET("/all", middleware.AbacMiddleware("services", "list"), sc.serviceController.ListServices)

	router.PUT("/client/update/:profileID/:serviceID", sc.serviceController.UpdateClientUserReviewOnProfile)
	router.PUT("/client/:profileID/:serviceID", sc.serviceController.HideProfileOwnerReview)

	router.PUT("/host/update", sc.serviceController.UpdateProfileOwnerReviewOnClientUser)
	router.PUT("/host/:profileID/:serviceID", sc.serviceController.HideUserReview)

}

func (sc *ServiceRouteController) ReviewsRoute(rg *gin.RouterGroup) {
	router := rg.Group("reviews")

	router.Use(middleware.DeserializeUser())

	router.PUT("/client", sc.serviceController.UpdateClientUserReviewOnProfile)
	router.PUT("/client/hide", sc.serviceController.HideProfileOwnerReview)

	router.PUT("/host", sc.serviceController.UpdateProfileOwnerReviewOnClientUser)
	router.PUT("/host/hide", sc.serviceController.HideUserReview)

}
