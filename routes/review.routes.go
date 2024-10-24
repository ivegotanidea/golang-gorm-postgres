package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/controllers"
	"github.com/ivegotanidea/golang-gorm-postgres/middleware"
)

type ReviewsRouteController struct {
	serviceController controllers.ServiceController
}

func NewRouteReviewController(serviceController controllers.ServiceController) ReviewsRouteController {
	return ReviewsRouteController{serviceController}
}

// @BasePath /api/v1/review

func (sc *ReviewsRouteController) ReviewsRoute(rg *gin.RouterGroup) {
	router := rg.Group("reviews")

	router.Use(middleware.DeserializeUser())

	router.PUT("/client", sc.serviceController.UpdateClientUserReviewOnProfile)
	router.PUT("/client/set-visibility", middleware.AbacMiddleware("reviews", "set-visibility"), sc.serviceController.HideProfileOwnerReview)

	router.PUT("/host", sc.serviceController.UpdateProfileOwnerReviewOnClientUser)
	router.PUT("/host/set-visibility", middleware.AbacMiddleware("reviews", "set-visibility"), sc.serviceController.HideUserReview)

}
