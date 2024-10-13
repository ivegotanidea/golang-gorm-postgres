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

// @BasePath /api/v1/service

func (sc *ServiceRouteController) ServiceRoute(rg *gin.RouterGroup) {
	router := rg.Group("services")

	router.Use(middleware.DeserializeUser())

	router.POST("/", sc.serviceController.CreateService)

	router.GET("/:profileID", sc.serviceController.GetProfileServices)

	router.GET("/:profileID/service/:serviceID", sc.serviceController.GetService)

	router.GET("/all", middleware.AbacMiddleware("services", "list"), sc.serviceController.ListServices)
}
