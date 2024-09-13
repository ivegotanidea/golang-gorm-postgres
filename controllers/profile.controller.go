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
		UserID:       currentUser.ID,
		CreatedAt:    now,
		UpdatedAt:    now,
		UpdatedBy:    currentUser.ID,
		CityID:       payload.CityID, // required, no need to check
		Active:       true,
		Phone:        payload.Phone,        // required, no need to check
		Name:         payload.Name,         // required, no need to check
		Age:          payload.Age,          // required, no need to check
		Height:       payload.Height,       // required, no need to check
		Weight:       payload.Weight,       // required, no need to check
		Bust:         payload.Bust,         // required, no need to check
		Bio:          payload.Bio,          // required, no need to check
		ContactPhone: payload.ContactPhone, // required, no need to check
		ContactTG:    payload.ContactTG,    // required, no need to check
	}

	// Process optional fields (omitempty)
	if payload.EthnosID != nil {
		newProfile.EthnosID = payload.EthnosID
	}
	if payload.HairColorID != nil {
		newProfile.HairColorID = payload.HairColorID
	}
	if payload.BodyTypeID != nil {
		newProfile.BodyTypeID = payload.BodyTypeID
	}
	if payload.IntimateHairCutID != nil {
		newProfile.IntimateHairCutID = payload.IntimateHairCutID
	}
	if payload.AddressLatitude != "" {
		newProfile.AddressLatitude = payload.AddressLatitude
	}
	if payload.AddressLongitude != "" {
		newProfile.AddressLongitude = payload.AddressLongitude
	}
	if payload.PriceInHouseContact != nil {
		newProfile.PriceInHouseContact = payload.PriceInHouseContact
	}
	if payload.PriceInHouseHour != nil {
		newProfile.PriceInHouseHour = payload.PriceInHouseHour
	}
	if payload.PriceSaunaContact != nil {
		newProfile.PriceSaunaContact = payload.PriceSaunaContact
	}
	if payload.PriceSaunaHour != nil {
		newProfile.PriceSaunaHour = payload.PriceSaunaHour
	}
	if payload.PriceVisitContact != nil {
		newProfile.PriceVisitContact = payload.PriceVisitContact
	}
	if payload.PriceVisitHour != nil {
		newProfile.PriceVisitHour = payload.PriceVisitHour
	}
	if payload.PriceCarContact != nil {
		newProfile.PriceCarContact = payload.PriceCarContact
	}
	if payload.PriceCarHour != nil {
		newProfile.PriceCarHour = payload.PriceCarHour
	}
	if payload.ContactWA != "" {
		newProfile.ContactWA = payload.ContactWA
	}

	// Insert profile into the database
	if err := tx.Create(&newProfile).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create profile"})
		return
	}

	var bodyArts []models.ProfileBodyArt

	// Insert associated body arts
	if len(payload.BodyArts) > 0 {
		for _, bodyArtReq := range payload.BodyArts {
			profileBodyArt := models.ProfileBodyArt{
				BodyArtID: bodyArtReq.ID,
				ProfileID: newProfile.ID,
			}
			bodyArts = append(bodyArts, profileBodyArt)
		}
		if err := tx.Create(&bodyArts).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create body arts connection" + err.Error()})
			return
		}
	}

	newProfile.BodyArts = bodyArts

	var photos []models.Photo

	// Insert associated photos
	if len(payload.Photos) > 0 {
		for _, photoReq := range payload.Photos {
			photo := models.Photo{
				ProfileID: newProfile.ID,
				URL:       photoReq.URL,
				CreatedAt: time.Now(),
			}
			photos = append(photos, photo)
		}
		if err := tx.Create(&photos).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create photos"})
			return
		}
	}

	newProfile.Photos = photos

	var options []models.ProfileOption

	// Insert associated profile options
	if len(payload.Options) > 0 {
		for _, optionReq := range payload.Options {
			option := models.ProfileOption{
				ProfileID:    newProfile.ID,
				ProfileTagID: int(optionReq.ProfileTagID),
				Price:        optionReq.Price,
				Comment:      optionReq.Comment,
			}
			options = append(options, option)
		}
		if err := tx.Create(&options).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create profile options" + err.Error()})
			return
		}
	}

	newProfile.ProfileOptions = options

	// Commit the transaction if everything was successful
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to commit transaction"})
		return
	}

	// Return the created profile in the response
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}

func (pc *ProfileController) UpdateOwnProfile(ctx *gin.Context) {
	profileId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload models.UpdateOwnProfileRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Find the existing profile
	var existingProfile models.Profile
	result := pc.DB.First(&existingProfile, "id = ?", profileId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No profile with that ID exists"})
		return
	}

	now := time.Now()

	// Start a transaction
	tx := pc.DB.Begin()

	// Only update fields that are not nil (omitempty)
	updatedProfile := models.Profile{
		UpdatedAt: now,
		UpdatedBy: currentUser.ID,
	}

	if payload.Active != nil &&
		existingProfile.Active != *payload.Active {

		updatedProfile.Active = *payload.Active
	}

	// Update profile in the database
	if err := tx.Model(&existingProfile).Select("Active").Updates(&updatedProfile).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update profile"})
		return
	}

	if payload.CityID != nil &&
		existingProfile.CityID != *payload.CityID {
		updatedProfile.CityID = *payload.CityID
	}

	if payload.Phone != "" && existingProfile.Phone != payload.Phone {
		updatedProfile.Phone = payload.Phone
	}

	if payload.Name != "" && existingProfile.Name != payload.Name {
		updatedProfile.Name = payload.Name
	}
	if payload.Age != nil &&
		existingProfile.Age != *payload.Age {
		updatedProfile.Age = *payload.Age
	}

	if payload.Height != nil &&
		existingProfile.Height != *payload.Height {
		updatedProfile.Height = *payload.Height
	}

	if payload.Weight != nil &&
		existingProfile.Weight != *payload.Weight {
		updatedProfile.Weight = *payload.Weight
	}

	if payload.Bust != nil &&
		existingProfile.Bust != *payload.Bust {
		updatedProfile.Bust = *payload.Bust
	}

	if payload.BodyTypeID != nil &&
		existingProfile.BodyTypeID != payload.BodyTypeID {
		updatedProfile.BodyTypeID = payload.BodyTypeID
	}

	if payload.EthnosID != nil &&
		existingProfile.EthnosID != payload.EthnosID {
		updatedProfile.EthnosID = payload.EthnosID
	}

	if payload.HairColorID != nil &&
		existingProfile.HairColorID != payload.HairColorID {
		updatedProfile.HairColorID = payload.HairColorID
	}

	if payload.IntimateHairCutID != nil &&
		existingProfile.IntimateHairCutID != payload.IntimateHairCutID {
		updatedProfile.IntimateHairCutID = payload.IntimateHairCutID
	}

	if payload.Bio != "" && existingProfile.Bio != payload.Bio {
		updatedProfile.Bio = payload.Bio
	}

	if payload.AddressLatitude != "" &&
		existingProfile.AddressLatitude != payload.AddressLatitude {
		updatedProfile.AddressLatitude = payload.AddressLatitude
	}
	if payload.AddressLongitude != "" &&
		existingProfile.AddressLongitude != payload.AddressLongitude {
		updatedProfile.AddressLongitude = payload.AddressLongitude
	}

	if payload.PriceInHouseNightRatio != nil &&
		existingProfile.PriceInHouseNightRatio != *payload.PriceInHouseNightRatio {

		updatedProfile.PriceInHouseNightRatio = *payload.PriceInHouseNightRatio
	}
	if payload.PriceInHouseContact != nil &&
		existingProfile.PriceInHouseContact != payload.PriceInHouseContact {

		updatedProfile.PriceInHouseContact = payload.PriceInHouseContact
	}
	if payload.PriceInHouseHour != nil &&
		existingProfile.PriceInHouseHour != payload.PriceInHouseHour {

		updatedProfile.PriceInHouseHour = payload.PriceInHouseHour
	}
	if payload.PrinceSaunaNightRatio != nil &&
		existingProfile.PrinceSaunaNightRatio != *payload.PrinceSaunaNightRatio {

		updatedProfile.PrinceSaunaNightRatio = *payload.PrinceSaunaNightRatio
	}
	if payload.PriceSaunaContact != nil &&
		existingProfile.PriceSaunaContact != payload.PriceSaunaContact {

		updatedProfile.PriceSaunaContact = payload.PriceSaunaContact
	}

	if payload.PriceSaunaHour != nil &&
		existingProfile.PriceSaunaHour != payload.PriceSaunaHour {
		updatedProfile.PriceSaunaHour = payload.PriceSaunaHour
	}

	if payload.PriceVisitNightRatio != nil &&
		existingProfile.PriceVisitNightRatio != *payload.PriceVisitNightRatio {

		updatedProfile.PriceVisitNightRatio = *payload.PriceVisitNightRatio
	}
	if payload.PriceVisitContact != nil &&
		existingProfile.PriceVisitContact != payload.PriceVisitContact {

		updatedProfile.PriceVisitContact = payload.PriceVisitContact
	}
	if payload.PriceVisitHour != nil &&
		existingProfile.PriceVisitHour != payload.PriceVisitHour {
		updatedProfile.PriceVisitHour = payload.PriceVisitHour
	}
	if payload.PriceCarNightRatio != nil &&
		existingProfile.PriceCarNightRatio != *payload.PriceCarNightRatio {
		updatedProfile.PriceCarNightRatio = *payload.PriceCarNightRatio
	}
	if payload.PriceCarContact != nil &&
		existingProfile.PriceCarContact != payload.PriceCarContact {
		updatedProfile.PriceCarContact = payload.PriceCarContact
	}
	if payload.PriceCarHour != nil &&
		existingProfile.PriceCarHour != payload.PriceCarHour {
		updatedProfile.PriceCarHour = payload.PriceCarHour
	}

	// Update profile in the database
	if err := tx.Model(&existingProfile).Updates(&updatedProfile).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update profile"})
		return
	}

	// Handle the update of BodyArts
	if payload.BodyArts != nil {
		if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&models.ProfileBodyArt{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old body arts"})
			return
		}

		if len(payload.BodyArts) > 0 {
			var bodyArts []models.ProfileBodyArt
			for _, bodyArtReq := range payload.BodyArts {
				profileBodyArt := models.ProfileBodyArt{
					BodyArtID: bodyArtReq.ID,
					ProfileID: existingProfile.ID,
				}
				bodyArts = append(bodyArts, profileBodyArt)
			}

			if err := tx.Create(&bodyArts).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update body arts"})
				return
			}
		}
	}

	// Handle the update of Photos
	if payload.Photos != nil {
		if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&models.Photo{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old photos"})
			return
		}

		if len(payload.Photos) > 0 {
			var photos []models.Photo
			for _, photoReq := range payload.Photos {
				photo := models.Photo{
					ProfileID: existingProfile.ID,
					URL:       photoReq.URL,
					CreatedAt: time.Now(),
				}
				photos = append(photos, photo)
			}

			if err := tx.Create(&photos).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update photos"})
				return
			}
		}
	}

	// Handle the update of ProfileOptions
	if payload.Options != nil {
		if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&models.ProfileOption{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old profile options"})
			return
		}

		if len(payload.Options) > 0 {
			var options []models.ProfileOption
			for _, optionReq := range payload.Options {
				option := models.ProfileOption{
					ProfileID:    existingProfile.ID,
					ProfileTagID: int(optionReq.ProfileTagID),
					Price:        optionReq.Price,
					Comment:      optionReq.Comment,
				}
				options = append(options, option)
			}

			if err := tx.Create(&options).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update profile options"})
				return
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to commit transaction"})
		return
	}

	// Return the updated profile
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": existingProfile})
}

func (pc *ProfileController) UpdateProfile(ctx *gin.Context) {
	profileId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload models.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Find the existing profile
	var existingProfile models.Profile
	result := pc.DB.First(&existingProfile, "id = ?", profileId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No profile with that ID exists"})
		return
	}

	now := time.Now()

	// Update the profile fields
	updatedProfile := models.Profile{
		Active:    payload.Active,
		Name:      payload.Name,
		Bio:       payload.Bio,
		Verified:  payload.Verified,
		Moderated: payload.Moderated,
		UpdatedAt: now,
		UpdatedBy: currentUser.ID,
	}

	tx := pc.DB.Begin()

	// Update profile in the database
	if err := tx.Model(&existingProfile).Updates(&updatedProfile).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update profile"})
		return
	}

	// Update associated body arts
	if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&models.ProfileBodyArt{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old body arts"})
		return
	}

	// Update associated photos
	if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&models.Photo{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old photos"})
		return
	}

	var photos []models.Photo
	for _, photoReq := range payload.Photos {
		photo := models.Photo{
			ProfileID: existingProfile.ID,
			URL:       photoReq.URL,
			CreatedAt: time.Now(),
		}
		photos = append(photos, photo)
	}

	// Batch insert photos
	if err := tx.Create(&photos).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update photos"})
		return
	}

	// Update associated profile options
	if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&models.ProfileOption{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old profile options"})
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to commit transaction"})
		return
	}

	// Return the updated profile
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

func (pc *ProfileController) ListProfiles(ctx *gin.Context) {
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

func (pc *ProfileController) GetMyProfiles(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	currentUser := ctx.MustGet("currentUser").(models.User)

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []models.Profile
	dbQuery := pc.DB.Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&profiles, "user_id = ?", currentUser.ID.String())

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(profiles), "data": profiles})
}

func (pc *ProfileController) FindProfiles(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	// Bind query parameters to the struct
	var query models.FindProfilesQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var profiles []models.Profile
	dbQuery := pc.DB.Limit(intLimit).Offset(offset)

	// Apply filtering based on query parameters
	if query.BodyTypeId != 0 {
		dbQuery = dbQuery.Where("body_type_id = ?", query.BodyTypeId)
	}
	if query.EthnosId != 0 {
		dbQuery = dbQuery.Where("ethnos_id = ?", query.EthnosId)
	}
	if query.HairColorId != 0 {
		dbQuery = dbQuery.Where("hair_color_id = ?", query.HairColorId)
	}
	if query.IntimateHairCutId != 0 {
		dbQuery = dbQuery.Where("intimate_hair_cut_id = ?", query.IntimateHairCutId)
	}
	if query.CityID != 0 {
		dbQuery = dbQuery.Where("city_id = ?", query.CityID)
	}
	if query.Active {
		dbQuery = dbQuery.Where("active = ?", query.Active)
	}
	if query.Phone != "" {
		dbQuery = dbQuery.Where("phone = ?", query.Phone)
	}
	if query.Age != 0 {
		dbQuery = dbQuery.Where("age = ?", query.Age)
	}
	if query.Name != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Height != 0 {
		dbQuery = dbQuery.Where("height = ?", query.Height)
	}
	if query.Weight != 0 {
		dbQuery = dbQuery.Where("weight = ?", query.Weight)
	}
	if query.Bust != 0 {
		dbQuery = dbQuery.Where("bust = ?", query.Bust)
	}
	if query.AddressLatitude != "" {
		dbQuery = dbQuery.Where("address_latitude = ?", query.AddressLatitude)
	}
	if query.AddressLongitude != "" {
		dbQuery = dbQuery.Where("address_longitude = ?", query.AddressLongitude)
	}
	if query.Moderated {
		dbQuery = dbQuery.Where("moderated = ?", query.Moderated)
	}
	if query.Verified {
		dbQuery = dbQuery.Where("verified = ?", query.Verified)
	}

	// Handle body arts and profile options if needed (complex relationships)
	if len(query.BodyArtIds) > 0 {
		dbQuery = dbQuery.Joins("JOIN profile_body_arts ON profiles.id = profile_body_arts.profile_id").
			Where("profile_body_arts.body_art_id IN ?", query.BodyArtIds)
	}
	if len(query.ProfileTagIds) > 0 {
		dbQuery = dbQuery.Joins("JOIN profile_options ON profiles.id = profile_options.profile_id").
			Where("profile_options.profile_tag_id IN ?", query.ProfileTagIds)
	}

	// Apply price range filters
	if query.PriceInHouseContactMin != 0 || query.PriceInHouseContactMax != 0 {
		dbQuery = dbQuery.Where("price_in_house_contact BETWEEN ? AND ?", query.PriceInHouseContactMin, query.PriceInHouseContactMax)
	}
	if query.PriceInHouseHourMin != 0 || query.PriceInHouseHourMax != 0 {
		dbQuery = dbQuery.Where("price_in_house_hour BETWEEN ? AND ?", query.PriceInHouseHourMin, query.PriceInHouseHourMax)
	}
	if query.PriceSaunaContactMin != 0 || query.PriceSaunaContactMax != 0 {
		dbQuery = dbQuery.Where("price_sauna_contact BETWEEN ? AND ?", query.PriceSaunaContactMin, query.PriceSaunaContactMax)
	}
	if query.PriceSaunaHourMin != 0 || query.PriceSaunaHourMax != 0 {
		dbQuery = dbQuery.Where("price_sauna_hour BETWEEN ? AND ?", query.PriceSaunaHourMin, query.PriceSaunaHourMax)
	}
	if query.PriceVisitContactMin != 0 || query.PriceVisitContactMax != 0 {
		dbQuery = dbQuery.Where("price_visit_contact BETWEEN ? AND ?", query.PriceVisitContactMin, query.PriceVisitContactMax)
	}
	if query.PriceVisitHourMin != 0 || query.PriceVisitHourMax != 0 {
		dbQuery = dbQuery.Where("price_visit_hour BETWEEN ? AND ?", query.PriceVisitHourMin, query.PriceVisitHourMax)
	}
	if query.PriceCarContactMin != 0 || query.PriceCarContactMax != 0 {
		dbQuery = dbQuery.Where("price_car_contact BETWEEN ? AND ?", query.PriceCarContactMin, query.PriceCarContactMax)
	}
	if query.PriceCarHourMin != 0 || query.PriceCarHourMax != 0 {
		dbQuery = dbQuery.Where("price_car_hour BETWEEN ? AND ?", query.PriceCarHourMin, query.PriceCarHourMax)
	}

	// Execute the query
	results := dbQuery.Find(&profiles)
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
