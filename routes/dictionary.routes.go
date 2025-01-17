package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivegotanidea/golang-gorm-postgres/controllers"
	"github.com/ivegotanidea/golang-gorm-postgres/middleware"
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

	// CRUD
	router.GET("/", dc.dictionaryController.ListDict)

	router.POST("/", middleware.DeserializeUser(), middleware.AbacMiddleware("dicts", "add"), dc.dictionaryController.CreateDict)
	router.PUT("/", middleware.AbacMiddleware("dicts", "update"), nil)
	router.DELETE("/", middleware.AbacMiddleware("dicts", "delete"), nil)

	// old style
	router.GET("/cities", dc.dictionaryController.ListCities)
	router.GET("/ethnos", dc.dictionaryController.ListEthnos)
	router.GET("/bodies", dc.dictionaryController.ListBodyTypes)
	router.GET("/arts", dc.dictionaryController.ListBodyArts)
	router.GET("/colors", dc.dictionaryController.ListHairColors)
	router.GET("/cuts", dc.dictionaryController.ListIntimateHairCuts)
	router.GET("/user/tags", dc.dictionaryController.ListUserTags)
	router.GET("/profile/tags", dc.dictionaryController.ListProfileTags)
}
