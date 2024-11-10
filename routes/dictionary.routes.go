package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/controllers"
)

type DictionaryRouteController struct {
	dictionaryController controllers.DictionaryController
}

func NewRouteDictionaryController(dictionaryController controllers.DictionaryController) DictionaryRouteController {
	return DictionaryRouteController{dictionaryController}
}

// @BasePath /api/v1/dictionary

func (dc *DictionaryRouteController) DictionaryRoute(rg *gin.RouterGroup) {

	router := rg.Group("dict")

	router.GET("/cities", dc.dictionaryController.ListCities)
}
