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

func (uc *ServiceRouteController) ServiceRoute(rg *gin.RouterGroup) {
	router := rg.Group("services")

	router.Use(middleware.DeserializeUser())
	
}
