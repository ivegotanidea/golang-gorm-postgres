package controllers

import (
	"fmt"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

const defaultUserAvatar = ""

// BotSignUpUser godoc
// @Summary Registers a new user via bot
// @Description Registers a new user by accepting Telegram user ID and other basic details. Automatically generates password.
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.BotSignUpInput true "Bot Signup Input"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /auth/bot/signup [post]
func (ac *AuthController) BotSignUpUser(ctx *gin.Context) {
	var payload *models.BotSignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	random := rand.New(rand.NewPCG(1, uint64(time.Now().Nanosecond())))
	generatedPassword := utils.GenerateRandomStringWithPrefix(random, 13, "")
	hashedPassword, err := utils.HashPassword(generatedPassword)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	telegramUserId, err := strconv.ParseInt(payload.TelegramUserId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{Status: "fail", Message: err.Error()})
	}

	now := time.Now()
	newUser := models.User{
		Name:           payload.Name,
		Phone:          payload.Phone,
		Password:       hashedPassword,
		TelegramUserId: telegramUserId,
		Verified:       false,
		HasProfile:     false,
		Avatar:         defaultUserAvatar,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, models.ErrorResponse{Status: "fail", Message: "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{Status: "fail", Message: "Something bad happened"})
		return
	}

	userResponse := &models.UserResponse{
		ID:             newUser.ID,
		TelegramUserID: newUser.TelegramUserId,
		Name:           newUser.Name,
		Phone:          newUser.Phone,
		Password:       generatedPassword,
		Avatar:         newUser.Avatar,
		Verified:       newUser.Verified,
		CreatedAt:      newUser.CreatedAt,
		UpdatedAt:      newUser.UpdatedAt,
		Tier:           newUser.Tier,
		Role:           newUser.Role,
	}
	ctx.JSON(http.StatusCreated, models.SuccessResponse{Status: "success", Data: userResponse})
}

// SignUpUser godoc
// @Summary Registers a new user
// @Description Registers a new user by accepting basic details and password confirmation.
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.SignUpInput true "SignUp Input"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Failure 502 {object} models.ErrorResponse
// @Router /auth/signup [post]
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:           "",
		Phone:          payload.Phone,
		Password:       hashedPassword,
		TelegramUserId: -1,
		Verified:       false,
		HasProfile:     false,
		Avatar:         "",
		CreatedAt:      now,
		UpdatedAt:      now,
		Tier:           "basic",
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, models.ErrorResponse{Status: "fail", Message: "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, models.ErrorResponse{Status: "fail", Message: "Something bad happened"})
		return
	}

	userResponse := &models.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Phone:     newUser.Phone,
		Avatar:    newUser.Avatar,
		Verified:  newUser.Verified,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		Tier:      newUser.Tier,
	}
	ctx.JSON(http.StatusCreated, models.SuccessResponse{Status: "success", Data: userResponse})
}

// BotSignInUser godoc
// @Summary Logs in a bot user
// @Description Authenticates a bot user by accepting Telegram User ID.
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.BotSignInInput true "Bot SignIn Input"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/bot/signin [post]
func (ac *AuthController) BotSignInUser(ctx *gin.Context) {
	var payload *models.BotSignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "telegram_user_id = ?", strings.ToLower(payload.TelegramUserId))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: "Invalid phone or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, models.TokenResponse{Status: "success", AccessToken: access_token})
}

// SignInUser godoc
// @Summary Logs in a user
// @Description Authenticates a user by accepting phone and password.
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.SignInInput true "SignIn Input"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/signin [post]
func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "phone = ?", strings.ToLower(payload.Phone))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: "Invalid phone or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: "Invalid email or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, models.TokenResponse{Status: "success", AccessToken: access_token})
}

// RefreshAccessToken godoc
// @Summary Refreshes access token
// @Description Refreshes the access token using the refresh token cookie.
// @Tags Auth
// @Produce json
// @Success 200 {object} models.TokenResponse
// @Failure 403 {object} models.ErrorResponse
// @Router /auth/refresh [post]
func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{Status: "fail", Message: message})
		return
	}

	config, _ := initializers.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{Status: "fail", Message: "user not exist"})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, models.ErrorResponse{Status: "fail", Message: err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, models.TokenResponse{Status: "success", AccessToken: access_token})
}

// LogoutUser godoc
// @Summary Logs out a user
// @Description Clears the access and refresh tokens and logs out the user.
// @Tags Auth
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Router /auth/logout [post]
func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, models.SuccessResponse{Status: "success"})
}
