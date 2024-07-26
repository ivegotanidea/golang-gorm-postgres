package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:         currentUser.ID,
		Name:       currentUser.Name,
		Phone:      currentUser.Phone,
		Avatar:     currentUser.Avatar,
		Verified:   currentUser.Verified,
		HasProfile: currentUser.HasProfile,
		CreatedAt:  currentUser.CreatedAt,
		UpdatedAt:  currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
