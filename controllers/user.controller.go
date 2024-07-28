package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
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
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Phone:     currentUser.Phone,
		Avatar:    currentUser.Avatar,
		Verified:  currentUser.Verified,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// todo available only for moderators / admins
func (uc *UserController) GetUsers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var users []models.User
	results := uc.DB.Limit(intLimit).Offset(offset).Find(&users)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}

func (uc *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var user models.User
	result := uc.DB.First(&user, "id = ?", userId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that id exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (uc *UserController) FindByTelegramId(ctx *gin.Context) {
	userId := ctx.Param("telegramUserId")

	var user models.User
	result := uc.DB.First(&user, "TelegramUserId = ?", userId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that id exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (uc *UserController) FindByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")

	var user models.User
	result := uc.DB.First(&user, "Phone = ?", phone)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that phone exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

// todo only available for moderators / admins
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	if currentUser.ID.String() == userId {
		fmt.Printf("User %v has commited self-deletion", currentUser.ID)
	}

	result := uc.DB.Delete(&models.Post{}, "id = ?", userId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// todo only available for moderators / admins
func (uc *UserController) DeleteUserByTelegramId(ctx *gin.Context) {
	userId := ctx.Param("telegramUserId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	if strconv.FormatInt(currentUser.TelegramUserId, 10) == userId {
		fmt.Printf("User %v has commited self-deletion", currentUser.ID)
	}

	result := uc.DB.Delete(&models.Post{}, "telegramUserId = ?", userId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var payload *models.UpdateUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedPost models.Post
	result := uc.DB.First(&updatedPost, "id = ?", userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	now := time.Now()

	type UpdateUser struct {
		Name      string    `json:"name,omitempty"`
		Phone     string    `json:"phone,omitempty"`
		Password  string    `json:"password,omitempty"`
		Avatar    string    `json:"photo,omitempty"`
		Verified  bool      `json:"verified,omitempty"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	userToUpdate := models.User{
		Name:      payload.Name,
		Phone:     payload.Phone,
		Password:  payload.Password,
		Avatar:    payload.Avatar,
		Verified:  payload.Verified,
		UpdatedAt: now,
	}

	uc.DB.Model(&updatedPost).Updates(userToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})
}

func (uc *UserController) UpdateUserByTelegramId(ctx *gin.Context) {
	userId := ctx.Param("telegramUserId")

	var payload *models.UpdateUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedPost models.Post
	result := uc.DB.First(&updatedPost, "telegramUserId = ?", userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	now := time.Now()

	type UpdateUser struct {
		Name      string    `json:"name,omitempty"`
		Phone     string    `json:"phone,omitempty"`
		Password  string    `json:"password,omitempty"`
		Avatar    string    `json:"photo,omitempty"`
		Verified  bool      `json:"verified,omitempty"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	userToUpdate := models.User{
		Name:      payload.Name,
		Phone:     payload.Phone,
		Password:  payload.Password,
		Avatar:    payload.Avatar,
		Verified:  payload.Verified,
		UpdatedAt: now,
	}

	uc.DB.Model(&updatedPost).Updates(userToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})
}
