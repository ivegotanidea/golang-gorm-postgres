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

	router.Use(middleware.DeserializeUser())

	router.GET("/", middleware.AbacMiddleware("users", "list"), uc.userController.FindUsers)

	router.GET("/me", uc.userController.GetMe)
	router.GET("/user", uc.userController.GetUser)

	router.DELETE("/user", uc.userController.DeleteSelf)
	router.DELETE("/user/:id", middleware.AbacMiddleware("users", "delete"), uc.userController.DeleteUser)

	router.PUT("/user", uc.userController.UpdateSelf)
	router.PUT("/user/:id", middleware.AbacMiddleware("users", "update"), uc.userController.UpdateUser)

	router.PUT("/role", middleware.AbacMiddleware("users", "promote"), uc.userController.AssignRole)
}
