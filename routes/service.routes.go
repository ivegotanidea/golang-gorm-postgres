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

	router.GET("/:profileID", middleware.AbacMiddleware("services", "get"), sc.serviceController.GetProfileServices)

	router.PUT("/:profileID", sc.serviceController.UpdateService)

	router.GET("/all", middleware.AbacMiddleware("services", "list"), sc.serviceController.ListServices)
}
