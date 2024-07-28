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
	router.GET("/:userId", middleware.DeserializeUser(), uc.userController.FindById)
	router.GET("/:telegramUserId", middleware.DeserializeUser(), uc.userController.FindByTelegramId)
	router.GET("/:phone", middleware.DeserializeUser(), uc.userController.FindByPhone)
	router.DELETE("/:userId", uc.userController.DeleteUser)
	router.DELETE("/:telegramUserId", uc.userController.DeleteUserByTelegramId)
	router.PUT("/:userId", uc.userController.UpdateUser)
	router.PUT("/:telegramUserId", uc.userController.UpdateUserByTelegramId)
}
