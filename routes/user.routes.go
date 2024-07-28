package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("users")

	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
	router.GET("/", middleware.DeserializeUser(), uc.userController.GetUsers)
	router.GET("/user", middleware.DeserializeUser(), uc.userController.FindUser)
	router.DELETE("/user", uc.userController.DeleteUser)
	router.PUT("/user", uc.userController.UpdateUser)
}
