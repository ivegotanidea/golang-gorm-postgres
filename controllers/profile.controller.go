package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type ProfileController struct {
	DB *gorm.DB
}

func NewProfileController(DB *gorm.DB) ProfileController {
	return ProfileController{DB}
}

func (pc *ProfileController) CreateProfile(ctx *gin.Context) {
	// Get the current user
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateProfileRequest

	// Bind and validate the input payload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Start a transaction
	tx := pc.DB.Begin()

	// Create the profile
	now := time.Now()
	newProfile := models.Profile{
		UserID:              currentUser.ID,
		CreatedAt:           now,
		UpdatedAt:           now,
		City:                payload.City,
		Active:              payload.Active,
		Phone:               payload.Phone,
		Name:                payload.Name,
		Age:                 payload.Age,
		Height:              payload.Height,
		Weight:              payload.Weight,
		Bust:                payload.Bust,
		Type:                payload.Type,
		Ethnos:              payload.Ethnos,
		Bio:                 payload.Bio,
		AddressLatitude:     payload.AddressLatitude,
		AddressLongitude:    payload.AddressLongitude,
		PriceInHouseContact: payload.PriceInHouseContact,
		PriceInHouseHour:    payload.PriceInHouseHour,
		PriceSaunaContact:   payload.PriceSaunaContact,
		PriceSaunaHour:      payload.PriceSaunaHour,
		PriceVisitContact:   payload.PriceVisitContact,
		PriceVisitHour:      payload.PriceVisitHour,
		PriceCarContact:     payload.PriceCarContact,
		PriceCarHour:        payload.PriceCarHour,
		ContactPhone:        payload.ContactPhone,
		ContactWA:           payload.ContactWA,
		ContactTG:           payload.ContactTG,
		Moderated:           false,
		Verified:            false,
	}

	// Insert profile into the database
	if err := tx.Create(&newProfile).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create profile"})
		return
	}

	// Insert associated photos
	var photos []models.Photo
	for _, photoReq := range payload.Photos {
		photo := models.Photo{
			ProfileID: newProfile.ID,
			URL:       photoReq.URL,
			CreatedAt: time.Now(),
		}
		photos = append(photos, photo)
	}

	// Batch insert photos
	if err := tx.Create(&photos).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create photos"})
		return
	}

	// Insert associated profile options
	var options []models.ProfileOption
	for _, optionReq := range payload.Options {
		option := models.ProfileOption{
			ProfileID:    newProfile.ID,
			ProfileTagID: optionReq.ProfileTagID,
			Price:        optionReq.Price,
		}
		options = append(options, option)
	}

	// Batch insert options
	if err := tx.Create(&options).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create profile options"})
		return
	}

	// Commit the transaction if everything was successful
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to commit transaction"})
		return
	}

	// Return the created profile in the response
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}

func (pc *ProfileController) UpdateProfile(ctx *gin.Context) {
	postId := ctx.Param("id")
	_ = ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedProfile models.Profile
	result := pc.DB.First(&updatedProfile, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	profileToUpdate := models.Profile{
		CreatedAt: updatedProfile.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedProfile).Updates(profileToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedProfile})
}

func (pc *ProfileController) FindProfileByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")

	var profile models.Profile
	result := pc.DB.First(&profile, "phone = ?", phone)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No profile with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": profile})
}

func (pc *ProfileController) FindProfiles(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []models.Profile
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&profiles)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(profiles), "data": profiles})
}

func (pc *ProfileController) DeleteProfile(ctx *gin.Context) {
	profileId := ctx.Param("id")

	result := pc.DB.Delete(&models.Profile{}, "id = ?", profileId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No profile with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
