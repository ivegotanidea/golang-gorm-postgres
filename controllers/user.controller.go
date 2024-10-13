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

// GetMe godoc
// @Summary Get current authenticated user
// @Description Retrieves the profile of the currently authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=models.UserResponse}
// @Failure 401 {object} models.ErrorResponse
// @Router /users/me [get]
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

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   userResponse,
	})
}

// FindUsers godoc
// @Summary Retrieve users based on the current user's role
// @Description Retrieves a paginated list of users based on the current user's role. Regular users can only see other users, while non-owners can see all users except owners.
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} models.SuccessResponse{data=[]models.User}
// @Failure 502 {object} models.ErrorResponse
// @Router /users [get]
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
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{
			Status:  "error",
			Message: results.Error.Error(),
		})
		return
	}

	intPage, _ = strconv.Atoi(page)
	intLimit, _ = strconv.Atoi(limit)

	ctx.JSON(http.StatusOK, models.SuccessPageResponse{
		Status:  "success",
		Results: len(users),
		Data:    users,
		Page:    intPage,
		Limit:   intLimit,
	})
}

// GetUser godoc
// @Summary Get a user by ID, Telegram user ID, or phone
// @Description Retrieve a user by providing their user ID, Telegram user ID, or phone number. At least one of these fields is required.
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string false "User ID"
// @Param telegramUserId query int false "Telegram User ID"
// @Param phone query string false "Phone number"
// @Success 200 {object} models.SuccessResponse{data=models.UserResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	var query *models.FindUserQuery

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	userId := query.Id
	telegramUserId := query.TelegramUserId
	phone := query.Phone

	if userId == "" && telegramUserId == 0 && phone == "" {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: "userId or telegramUserId or phone is required",
		})
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
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "fail",
				Message: "No user with that ID exists",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: result.Error.Error(),
		})
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

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   userResponse,
	})
}

// DeleteSelf godoc
// @Summary Delete the currently authenticated user
// @Description Allows the current user to delete their own account.
// @Tags Users
// @Produce json
// @Success 204 {object} nil
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/self [delete]
func (uc *UserController) DeleteSelf(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	userId := currentUser.ID.String()

	fmt.Printf("User %v has committed self-deletion", currentUser.ID)

	var result *gorm.DB

	if userId != "" {
		result = uc.DB.Delete(&models.User{}, "id = ?", userId)
	} else {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: "userId or telegramUserId or phone is required",
		})
		return
	}

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "fail",
			Message: "No user with that ID exists",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Allows an authorized user to delete another user by their ID, with role-based restrictions.
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} nil
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var targetUser models.User

	if err := initializers.DB.First(&targetUser, "id = ?", userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "fail",
			Message: "User not found",
		})
		return
	}

	// Role-based restrictions
	if currentUser.Role == "moderator" && (targetUser.Role == "moderator" || targetUser.Role == "admin" || targetUser.Role == "owner") {
		ctx.JSON(http.StatusForbidden, models.ErrorResponse{
			Status:  "fail",
			Message: "You are not authorized to delete this user",
		})
		return
	}

	if currentUser.Role == "admin" && (targetUser.Role == "admin" || targetUser.Role == "owner") {
		ctx.JSON(http.StatusForbidden, models.ErrorResponse{
			Status:  "fail",
			Message: "You are not authorized to delete this user",
		})
		return
	}

	// Proceed with deletion
	var result *gorm.DB
	if userId != "" {
		result = uc.DB.Delete(&models.User{}, "id = ?", userId)
	} else {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: "User ID is required",
		})
		return
	}

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "fail",
			Message: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UpdateSelf godoc
// @Summary Update the current user's information
// @Description Allows the current user to update their own profile, including name, phone, and avatar.
// @Tags Users
// @Accept json
// @Produce json
// @Param body body models.UpdateUser true "User Update Payload"
// @Success 200 {object} models.SuccessResponse{data=models.UserResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /users/me [put]
func (uc *UserController) UpdateSelf(ctx *gin.Context) {

	currentUser := ctx.MustGet("currentUser").(models.User)
	userId := currentUser.ID.String()

	var payload *models.UpdateUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	if err := uc.validator.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	var updatedUser models.User
	result := uc.DB.First(&updatedUser, "id = ?", userId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "fail",
			Message: "User not found",
		})
		return
	}

	// Validate and update avatar if needed
	avatarUrl, err := checkAvatar(payload.Avatar, updatedUser.Avatar)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: err,
		})
		return
	}

	now := time.Now()

	// Prepare the updated user data
	userToUpdate := models.User{
		Name:      payload.Name,
		Phone:     payload.Phone,
		Avatar:    avatarUrl,
		UpdatedAt: now,
	}

	// Apply the updates to the database
	uc.DB.Model(&updatedUser).Updates(userToUpdate)

	// Prepare the user response
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

	// Return the updated user data in the response
	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   userResponse,
	})
}

// UpdateUser godoc
// @Summary Update a user's information (privileged access)
// @Description Allows privileged users to update user details, including Telegram ID, verification status, tier, and active status.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body models.UpdateUserPrivileged true "User Update Payload"
// @Success 200 {object} models.SuccessResponse{data=models.User}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	// Parse and bind the payload from the request
	var payload *models.UpdateUserPrivileged
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Validate the payload
	if err := uc.validator.Struct(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	// Find the user to be updated
	var updatedUser models.User
	result := uc.DB.First(&updatedUser, "id = ?", userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "fail",
			Message: "User not found",
		})
		return
	}

	// Parse TelegramUserId if provided
	var newTelegramId int64
	var err error
	if payload.TelegramUserId != "" {
		newTelegramId, err = strconv.ParseInt(payload.TelegramUserId, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Status:  "fail",
				Message: "Invalid telegram id",
			})
			return
		}
	}

	// Set the current time
	now := time.Now()

	// Prepare the updated user data
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

	// Apply the updates to the database
	tx := uc.DB.Model(&updatedUser).Updates(userToUpdate)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "fail",
			Message: tx.Error.Error(),
		})
		return
	}

	// Return the updated user data in the response
	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status: "success",
		Data:   updatedUser,
	})
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
