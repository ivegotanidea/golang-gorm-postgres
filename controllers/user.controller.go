package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type UserController struct {
	DB        *gorm.DB
	validator *validator.Validate
}

func NewUserController(DB *gorm.DB) UserController {

	v := validator.New()
	err := v.RegisterValidation("imageurl", utils.ValidateImageURL)

	if err != nil {
		panic(err)
	}

	return UserController{DB, v}
}

func checkAvatar(newAvatarUrl string, oldAvatarUrl string) (string, string) {

	if newAvatarUrl == "" {
		return defaultUserAvatar, ""
	}

	if newAvatarUrl == oldAvatarUrl && oldAvatarUrl != defaultUserAvatar {
		return newAvatarUrl, ""
	}

	if _, err := url.ParseRequestURI(newAvatarUrl); err != nil {
		return "", err.Error()
	}

	return newAvatarUrl, ""
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

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": userResponse})
}

func (uc *UserController) FindUsers(ctx *gin.Context) {

	currentUser := ctx.MustGet("currentUser").(models.User)

	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var users []models.User

	var results *gorm.DB

	if currentUser.Role == "user" {
		results = uc.DB.Limit(intLimit).Offset(offset).Find(&users, "role = ?", "user")
	} else {
		results = uc.DB.Limit(intLimit).Offset(offset).Find(&users, "role != ?", "owner")
	}

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users, "page": page, "limit": limit})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	var query *models.FindUserQuery

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userId := query.Id
	telegramUserId := query.TelegramUserId
	phone := query.Phone

	if userId == "" && telegramUserId == 0 && phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId or telegramUserId or phone is required"})
		return
	}

	var user models.User

	var result *gorm.DB

	if userId != "" {
		result = uc.DB.First(&user, "id = ?", userId)
	} else if telegramUserId != 0 {
		result = uc.DB.First(&user, "telegram_user_id = ?", telegramUserId)
	} else if phone != "" {
		result = uc.DB.First(&user, "phone = ?", phone)
	}

	if result.Error != nil {

		if result.Error.Error() == "record not found" && result.RowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that id exists"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		return
	}

	userResponse := &models.UserResponse{
		ID:             user.ID,
		TelegramUserID: user.TelegramUserId,
		Name:           user.Name,
		Phone:          user.Phone,
		Avatar:         user.Avatar,
		Verified:       user.Verified,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Tier:           user.Tier,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": userResponse})
}

func (uc *UserController) DeleteSelf(ctx *gin.Context) {

	currentUser := ctx.MustGet("currentUser").(models.User)
	userId := currentUser.ID.String()

	fmt.Printf("User %v has commited self-deletion", currentUser.ID)

	var result *gorm.DB

	if userId != "" {
		result = uc.DB.Delete(&models.User{}, "id = ?", userId)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId or telegramUserId or phone is required"})
		return
	}

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	currentUser := ctx.MustGet("currentUser").(models.User)

	var targetUser models.User

	if err := initializers.DB.First(&targetUser, "id = ?", userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	if currentUser.Role == "moderator" && (targetUser.Role == "moderator" || targetUser.Role == "admin" || targetUser.Role == "owner") {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	if currentUser.Role == "admin" && (targetUser.Role == "admin" || targetUser.Role == "owner") {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	var result *gorm.DB

	if userId != "" {
		result = uc.DB.Delete(&models.User{}, "id = ?", userId)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId or telegramUserId or phone is required"})
		return
	}

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (uc *UserController) UpdateSelf(ctx *gin.Context) {

	currentUser := ctx.MustGet("currentUser").(models.User)
	userId := currentUser.ID.String()

	var payload *models.UpdateUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := uc.validator.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedUser models.User

	result := uc.DB.First(&updatedUser, "id = ?", userId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	avatarUrl, err := checkAvatar(payload.Avatar, updatedUser.Avatar)

	if err != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err})
		return
	}

	now := time.Now()

	userToUpdate := models.User{
		Name:      payload.Name,
		Phone:     payload.Phone,
		Avatar:    avatarUrl,
		UpdatedAt: now,
	}

	uc.DB.Model(&updatedUser).Updates(userToUpdate)

	userResponse := &models.UserResponse{
		ID:             updatedUser.ID,
		TelegramUserID: updatedUser.TelegramUserId,
		Name:           updatedUser.Name,
		Phone:          updatedUser.Phone,
		Avatar:         updatedUser.Avatar,
		Verified:       updatedUser.Verified,
		CreatedAt:      updatedUser.CreatedAt,
		UpdatedAt:      updatedUser.UpdatedAt,
		Tier:           updatedUser.Tier,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": userResponse})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	var payload *models.UpdateUserPrivileged
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := uc.validator.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedUser models.User
	result := uc.DB.First(&updatedUser, "id = ?", userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	var newTelegramId int64
	var err error

	if payload.TelegramUserId != "" {

		newTelegramId, err = strconv.ParseInt(payload.TelegramUserId, 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid telegram id"})
			return
		}
	}

	now := time.Now()

	userToUpdate := models.User{
		Name:           payload.Name,
		Phone:          payload.Phone,
		Avatar:         payload.Avatar,
		TelegramUserId: newTelegramId,
		Verified:       payload.Verified,
		Tier:           payload.Tier,
		Active:         payload.Active,
		UpdatedAt:      now,
	}

	tx := uc.DB.Model(&updatedUser).Updates(userToUpdate)

	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": tx.Error.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
}

func (uc *UserController) AssignRole(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload models.AssignRole
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := uc.validator.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var targetUser models.User
	if err := initializers.DB.First(&targetUser, "id = ?", payload.Id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	if currentUser.Role == "admin" && (targetUser.Role == "admin" || targetUser.Role == "owner") {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail"})
		return
	}

	if targetUser.HasProfile {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "User has profile"})
		return
	}

	// Assign the role to the targetUser
	targetUser.Role = payload.Role

	// Automatically assign the "guru" tier if the role is "moderator"
	if payload.Role == "moderator" || payload.Role == "admin" {
		targetUser.Tier = "guru"
	}

	if err := initializers.DB.Save(&targetUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to change user's role"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": targetUser})

}
