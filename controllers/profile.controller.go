package controllers

import (
	"fmt"
	"github.com/ivegotanidea/golang-gorm-postgres/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type ProfileController struct {
	DB *gorm.DB
}

func NewProfileController(DB *gorm.DB) ProfileController {
	return ProfileController{DB}
}

// CreateProfile godoc
//
//	@Summary		Creates a new profile
//	@Description	Creates a new profile for the current user
//	@Tags			Profiles
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateProfileRequest	true	"Create Profile Request"
//	@Success		201		{object}	SuccessResponse[ProfileResponse]
//	@Failure		400		{object}	ErrorResponse
//	@Failure		403		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/profiles [post]
func (pc *ProfileController) CreateProfile(ctx *gin.Context) {
	// Get the current user
	currentUser := ctx.MustGet("currentUser").(User)

	if currentUser.Role != "user" {
		ctx.JSON(http.StatusForbidden, ErrorResponse{Status: "error", Message: "unauthorized"})
		return
	}

	var payload *CreateProfileRequest

	// Bind and validate the input payload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: err.Error()})
		return
	}

	// Start a transaction
	tx := pc.DB.Begin()

	// Create the profile
	now := time.Now()
	newProfile := Profile{
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
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to create profile: %s", err.Error())})
		return
	}

	var bodyArts []ProfileBodyArt

	// Insert associated body arts
	if len(payload.BodyArts) > 0 {
		for _, bodyArtReq := range payload.BodyArts {
			profileBodyArt := ProfileBodyArt{
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

	var photos []Photo

	// Insert associated photos
	if len(payload.Photos) > 0 {
		for _, photoReq := range payload.Photos {
			photo := Photo{
				ProfileID: newProfile.ID,
				URL:       photoReq.URL,
				CreatedAt: time.Now(),
			}
			photos = append(photos, photo)
		}
		if err := tx.Create(&photos).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to create photos: %s", err.Error())})
			return
		}
	}

	newProfile.Photos = photos

	var options []ProfileOption

	// Insert associated profile options
	if len(payload.Options) > 0 {
		for _, optionReq := range payload.Options {
			option := ProfileOption{
				ProfileID:    newProfile.ID,
				ProfileTagID: int(optionReq.ProfileTagID),
				Price:        optionReq.Price,
				Comment:      optionReq.Comment,
			}
			options = append(options, option)
		}
		if err := tx.Create(&options).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Failed to create profile options: %s", err.Error())})
			return
		}
	}

	newProfile.ProfileOptions = options

	profileResponse := utils.MapProfile(&newProfile)

	// Commit the transaction if everything was successful
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: err.Error()})
		return
	}

	// Return the created profile in the response
	ctx.JSON(http.StatusCreated, SuccessResponse[*ProfileResponse]{Status: "success", Data: profileResponse})
}

// UpdateOwnProfile godoc
//
//	@Summary		User updates his own profile
//	@Description	User updates his own profile with the provided fields
//	@Tags			Profiles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Profile ID"
//	@Param			body	body		UpdateOwnProfileRequest	true	"Profile Update Payload"
//	@Success		200		{object}	SuccessResponse[ProfileResponse]
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/profiles/my/{id} [put]
func (pc *ProfileController) UpdateOwnProfile(ctx *gin.Context) {
	profileId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(User)

	var payload UpdateOwnProfileRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: err.Error()})
		return
	}

	// Find the existing profile
	var existingProfile Profile
	result := pc.DB.First(&existingProfile, "id = ?", profileId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Status: "error", Message: "No profile with that ID exists"})
		return
	}

	now := time.Now()

	// Initialize the map to keep track of fields that need to be updated
	updateFields := map[string]interface{}{
		"UpdatedAt": now,
		"UpdatedBy": currentUser.ID,
	}

	// Check each field to see if it has changed and should be updated
	if payload.Active != nil && *payload.Active != existingProfile.Active {
		updateFields["Active"] = *payload.Active
	}

	if payload.CityID != nil && *payload.CityID != existingProfile.CityID {
		updateFields["CityID"] = *payload.CityID
	}

	if payload.Phone != "" && payload.Phone != existingProfile.Phone {
		updateFields["Phone"] = payload.Phone
	}

	if payload.Name != "" && payload.Name != existingProfile.Name {
		updateFields["Name"] = payload.Name
	}

	if payload.Age != nil && *payload.Age != existingProfile.Age {
		updateFields["Age"] = *payload.Age
	}

	if payload.Height != nil && *payload.Height != existingProfile.Height {
		updateFields["Height"] = *payload.Height
	}

	if payload.Weight != nil && *payload.Weight != existingProfile.Weight {
		updateFields["Weight"] = *payload.Weight
	}

	if payload.Bust != nil && *payload.Bust != existingProfile.Bust {
		updateFields["Bust"] = *payload.Bust
	}

	if payload.BodyTypeID != nil && payload.BodyTypeID != existingProfile.BodyTypeID {
		updateFields["BodyTypeID"] = payload.BodyTypeID
	}

	if payload.EthnosID != nil && payload.EthnosID != existingProfile.EthnosID {
		updateFields["EthnosID"] = payload.EthnosID
	}

	if payload.HairColorID != nil && payload.HairColorID != existingProfile.HairColorID {
		updateFields["HairColorID"] = payload.HairColorID
	}

	if payload.IntimateHairCutID != nil && payload.IntimateHairCutID != existingProfile.IntimateHairCutID {
		updateFields["IntimateHairCutID"] = payload.IntimateHairCutID
	}

	if payload.Bio != "" && payload.Bio != existingProfile.Bio {
		updateFields["Bio"] = payload.Bio
	}

	if payload.AddressLatitude != "" && payload.AddressLatitude != existingProfile.AddressLatitude {
		updateFields["AddressLatitude"] = payload.AddressLatitude
	}

	if payload.AddressLongitude != "" && payload.AddressLongitude != existingProfile.AddressLongitude {
		updateFields["AddressLongitude"] = payload.AddressLongitude
	}

	if payload.PriceInHouseNightRatio != nil && *payload.PriceInHouseNightRatio != existingProfile.PriceInHouseNightRatio {
		updateFields["PriceInHouseNightRatio"] = *payload.PriceInHouseNightRatio
	}

	if payload.PriceInHouseContact != nil && payload.PriceInHouseContact != existingProfile.PriceInHouseContact {
		updateFields["PriceInHouseContact"] = payload.PriceInHouseContact
	}

	if payload.PriceInHouseHour != nil && payload.PriceInHouseHour != existingProfile.PriceInHouseHour {
		updateFields["PriceInHouseHour"] = payload.PriceInHouseHour
	}

	if payload.PrinceSaunaNightRatio != nil && *payload.PrinceSaunaNightRatio != existingProfile.PrinceSaunaNightRatio {
		updateFields["PrinceSaunaNightRatio"] = *payload.PrinceSaunaNightRatio
	}

	if payload.PriceSaunaContact != nil && payload.PriceSaunaContact != existingProfile.PriceSaunaContact {
		updateFields["PriceSaunaContact"] = payload.PriceSaunaContact
	}

	if payload.PriceSaunaHour != nil && payload.PriceSaunaHour != existingProfile.PriceSaunaHour {
		updateFields["PriceSaunaHour"] = payload.PriceSaunaHour
	}

	if payload.PriceVisitNightRatio != nil && *payload.PriceVisitNightRatio != existingProfile.PriceVisitNightRatio {
		updateFields["PriceVisitNightRatio"] = *payload.PriceVisitNightRatio
	}

	if payload.PriceVisitContact != nil && payload.PriceVisitContact != existingProfile.PriceVisitContact {
		updateFields["PriceVisitContact"] = payload.PriceVisitContact
	}

	if payload.PriceVisitHour != nil && payload.PriceVisitHour != existingProfile.PriceVisitHour {
		updateFields["PriceVisitHour"] = payload.PriceVisitHour
	}

	if payload.PriceCarNightRatio != nil && *payload.PriceCarNightRatio != existingProfile.PriceCarNightRatio {
		updateFields["PriceCarNightRatio"] = *payload.PriceCarNightRatio
	}

	if payload.PriceCarContact != nil && payload.PriceCarContact != existingProfile.PriceCarContact {
		updateFields["PriceCarContact"] = payload.PriceCarContact
	}

	if payload.PriceCarHour != nil && payload.PriceCarHour != existingProfile.PriceCarHour {
		updateFields["PriceCarHour"] = payload.PriceCarHour
	}

	// Start a transaction
	tx := pc.DB.Begin()

	// Update only the fields that have changed
	if err := tx.Model(&existingProfile).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to update profile"})
		return
	}

	// Handle the update of BodyArts
	if payload.BodyArts != nil {
		if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&ProfileBodyArt{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to delete old body arts"})
			return
		}

		if len(payload.BodyArts) > 0 {
			var bodyArts []ProfileBodyArt
			for _, bodyArtReq := range payload.BodyArts {
				profileBodyArt := ProfileBodyArt{
					BodyArtID: bodyArtReq.ID,
					ProfileID: existingProfile.ID,
				}
				bodyArts = append(bodyArts, profileBodyArt)
			}

			if err := tx.Create(&bodyArts).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to update body arts"})
				return
			}
		}
	}

	// Handle the update of Photos
	if payload.Photos != nil {
		if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&Photo{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to delete old photos"})
			return
		}

		if len(payload.Photos) > 0 {
			var photos []Photo
			for _, photoReq := range payload.Photos {
				photo := Photo{
					ProfileID: existingProfile.ID,
					URL:       photoReq.URL,
					CreatedAt: time.Now(),
				}
				photos = append(photos, photo)
			}

			if err := tx.Create(&photos).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to update photos"})
				return
			}
		}
	}

	// Handle the update of ProfileOptions
	if payload.Options != nil {
		if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&ProfileOption{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to delete old profile options"})
			return
		}

		if len(payload.Options) > 0 {
			var options []ProfileOption
			for _, optionReq := range payload.Options {
				option := ProfileOption{
					ProfileID:    existingProfile.ID,
					ProfileTagID: int(optionReq.ProfileTagID),
					Price:        optionReq.Price,
					Comment:      optionReq.Comment,
				}
				options = append(options, option)
			}

			if err := tx.Create(&options).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to update profile options"})
				return
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: "Failed to commit transaction"})
		return
	}

	profileResponse := utils.MapProfile(&existingProfile)

	// Return the updated profile
	ctx.JSON(http.StatusOK, SuccessResponse[*ProfileResponse]{Status: "success", Data: profileResponse})
}

// UpdateProfile godoc
//
//	@Summary		Updates an existing profile
//	@Description	Updates the profile with the given ID, allows updating specific fields
//	@Tags			Profiles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Profile ID"
//	@Param			body	body		UpdateProfileRequest	true	"Profile Update Payload"
//	@Success		200		{object}	SuccessResponse[ProfileResponse]
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/profiles/update/{id} [put]
func (pc *ProfileController) UpdateProfile(ctx *gin.Context) {
	profileId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(User)

	var payload UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", Message: err.Error()})
		return
	}

	// Find the existing profile
	var existingProfile Profile
	result := pc.DB.First(&existingProfile, "id = ?", profileId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Status: "error", Message: "No profile with that ID exists"})
		return
	}

	now := time.Now()

	// Initialize the map to keep track of fields that need to be updated
	updateFields := map[string]interface{}{
		"UpdatedAt": now,
		"UpdatedBy": currentUser.ID,
	}

	// Check each field to see if it has changed and should be updated
	if payload.Active != nil && *payload.Active != existingProfile.Active {
		updateFields["Active"] = *payload.Active
	}

	if payload.Verified != nil {
		updateFields["Verified"] = *payload.Verified
		updateFields["VerifiedAt"] = now
		updateFields["VerifiedBy"] = currentUser.ID
	}

	if payload.Moderated != nil {
		updateFields["Moderated"] = *payload.Moderated
		updateFields["ModeratedAt"] = now
		updateFields["ModeratedBy"] = currentUser.ID
	}

	if payload.Name != "" && payload.Name != existingProfile.Name {
		updateFields["Name"] = payload.Name
	}

	if payload.Bio != "" && payload.Bio != existingProfile.Bio {
		updateFields["Bio"] = payload.Bio
	}

	// Start a transaction
	tx := pc.DB.Begin()

	// Update only the fields that have changed
	if err := tx.Model(&existingProfile).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Update failed: %s", err.Error())})
		return
	}

	// Handle the update of Photos
	if payload.Photos != nil {
		if err := tx.Where("profile_id = ?", existingProfile.ID).Delete(&Photo{}).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Update failed: %s", err.Error())})
			return
		}

		if len(payload.Photos) > 0 {
			var photos []Photo
			for _, photoReq := range payload.Photos {
				photo := Photo{
					ProfileID: existingProfile.ID,
					URL:       photoReq.URL,
					CreatedAt: time.Now(),
				}
				photos = append(photos, photo)
			}

			if err := tx.Create(&photos).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Update failed: %s", err.Error())})
				return
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Status: "error", Message: fmt.Sprintf("Update failed: %s", err.Error())})
		return
	}

	profileResponse := utils.MapProfile(&existingProfile)

	// Return the updated profile
	ctx.JSON(http.StatusOK, SuccessResponse[*ProfileResponse]{Status: "success", Data: profileResponse})
}

// FindProfileByID godoc
//
//	@Summary		Find a profile by ID
//	@Description	Retrieves a profile based on the id
//	@Tags			Profiles
//	@Produce		json
//	@Param			id	path		string	true	"Profile ID"
//	@Success		200		{object}	SuccessResponse[ProfileResponse]
//	@Failure		404		{object}	ErrorResponse
//	@Router			/profiles/{id} [get]
func (pc *ProfileController) FindProfileByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var profile Profile

	result := pc.DB.Preload("Photos").
		Preload("BodyArts").
		Preload("ProfileOptions.ProfileTag").
		First(&profile, "id = ?", id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Status: "error", Message: "No profile with that title exists"})
		return
	}

	profileResponse := utils.MapProfile(&profile)
	ctx.JSON(http.StatusOK, SuccessResponse[*ProfileResponse]{Status: "success", Data: profileResponse})
}

// FindProfileByPhone godoc
//
//	@Summary		Find a profile by phone number
//	@Description	Retrieves a profile based on the phone number provided
//	@Tags			Profiles
//	@Produce		json
//	@Param			phone	path		string	true	"Phone Number"
//	@Success		200		{object}	SuccessResponse[ProfileResponse]
//	@Failure		404		{object}	ErrorResponse
//	@Router			/profiles/{phone} [get]
func (pc *ProfileController) FindProfileByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")

	var profile Profile

	result := pc.DB.Preload("Photos").
		Preload("BodyArts").
		Preload("ProfileOptions.ProfileTag").
		First(&profile, "phone = ?", phone)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Status: "error", Message: "No profile with that title exists"})
		return
	}

	profileResponse := utils.MapProfile(&profile)
	ctx.JSON(http.StatusOK, SuccessResponse[*ProfileResponse]{Status: "success", Data: profileResponse})
}

// ListProfiles godoc
//
//	@Summary		Lists all profiles with pagination, auth required
//	@Description	Retrieves all profiles, supports pagination
//	@Tags			Profiles
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[ProfileResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/profiles/all [get]
func (pc *ProfileController) ListProfiles(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []Profile

	dbQuery := pc.DB.Preload("Photos").
		Preload("BodyArts").
		Preload("ProfileOptions.ProfileTag").
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&profiles)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	profileResponses := make([]ProfileResponse, len(profiles))
	for i, profile := range profiles {
		profileResponses[i] = *utils.MapProfile(&profile) // Assuming you have the mapProfile function
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]ProfileResponse]{
		Status:  "success",
		Data:    profileResponses,
		Results: len(profiles),
		Page:    intPage,
	})
}

// todo: ListProfilesNonAuth
// todo: restricting logic for non authorized users
// todo: add tests

// ListProfilesNonAuth godoc
//
//	@Summary		Lists all active profiles with pagination, no auth required
//	@Description	Retrieves all profiles, supports pagination
//	@Tags			Profiles
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[ProfileResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/profiles/list [get]
func (pc *ProfileController) ListProfilesNonAuth(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	var sex = ctx.DefaultQuery("sex", "female")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []Profile

	dbQuery := pc.DB.Preload("Photos").
		Preload("BodyArts").
		Preload("ProfileOptions.ProfileTag").
		Where("active = ?", true).
		Where("sex = ?", sex).
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&profiles)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	profileResponses := make([]ProfileResponse, len(profiles))
	for i, profile := range profiles {
		profileResponses[i] = *utils.MapProfile(&profile) // Assuming you have the mapProfile function
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]ProfileResponse]{
		Status:  "success",
		Data:    profileResponses,
		Results: len(profiles),
		Page:    intPage,
	})
}

// GetMyProfiles godoc
//
//	@Summary		Get current user's profiles
//	@Description	Retrieves the profiles created by the currently authenticated user
//	@Tags			Profiles
//	@Produce		json
//	@Param			page	query		string	false	"Page number"
//	@Param			limit	query		string	false	"Items per page"
//	@Success		200		{object}	SuccessPageResponse[ProfileResponse[]]
//	@Failure		502		{object}	ErrorResponse
//	@Router			/profiles/my [get]
func (pc *ProfileController) GetMyProfiles(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	currentUser := ctx.MustGet("currentUser").(User)

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []Profile

	dbQuery := pc.DB.Preload("Photos").
		Preload("BodyArts").
		Preload("ProfileOptions.ProfileTag").
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&profiles, "user_id = ?", currentUser.ID.String())

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{Status: "error", Message: results.Error.Error()})
		return
	}

	profileResponses := make([]ProfileResponse, len(profiles))
	for i, profile := range profiles {
		profileResponses[i] = *utils.MapProfile(&profile) // Assuming you have the mapProfile function
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse[[]ProfileResponse]{
		Status:  "success",
		Data:    profileResponses,
		Results: len(profiles),
		Page:    intPage,
	})
}

// FindProfiles godoc
//
//	@Summary		Search for profiles
//	@Description	Retrieves profiles based on filters provided in the query
//	@Tags			Profiles
//	@Accept			json
//	@Produce		json
//	@Param			body	body		FindProfilesQuery	true	"Search Filters"
//	@Success		200		{object}	SuccessPageResponse[ProfileResponse[]]
//	@Failure		400		{object}	ErrorResponse
//	@Failure		502		{object}	ErrorResponse
//	@Router			/profiles/search [post]
func (pc *ProfileController) FindProfiles(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	// Bind the JSON payload to the struct
	var query FindProfilesQuery
	if err := ctx.ShouldBindJSON(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	var profiles []Profile
	dbQuery := pc.DB.Limit(intLimit).Offset(offset)

	// Apply filtering based on query parameters
	if query.BodyTypeId != nil {
		dbQuery = dbQuery.Where("body_type_id = ?", query.BodyTypeId)
	}
	if query.EthnosId != nil {
		dbQuery = dbQuery.Where("ethnos_id = ?", query.EthnosId)
	}
	if query.HairColorId != nil {
		dbQuery = dbQuery.Where("hair_color_id = ?", query.HairColorId)
	}
	if query.IntimateHairCutId != nil {
		dbQuery = dbQuery.Where("intimate_hair_cut_id = ?", query.IntimateHairCutId)
	}
	if query.CityID != nil {
		dbQuery = dbQuery.Where("city_id = ?", query.CityID)
	}
	if query.Active != nil {
		dbQuery = dbQuery.Where("active = ?", query.Active)
	}
	if query.Phone != "" {
		dbQuery = dbQuery.Where("phone = ?", query.Phone)
	}
	if query.Age != nil {
		dbQuery = dbQuery.Where("age = ?", query.Age)
	}
	if query.Name != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Height != nil {
		dbQuery = dbQuery.Where("height = ?", query.Height)
	}
	if query.Weight != nil {
		dbQuery = dbQuery.Where("weight = ?", query.Weight)
	}
	if query.Bust != nil {
		dbQuery = dbQuery.Where("bust = ?", query.Bust)
	}
	if query.AddressLatitude != "" {
		dbQuery = dbQuery.Where("address_latitude = ?", query.AddressLatitude)
	}
	if query.AddressLongitude != "" {
		dbQuery = dbQuery.Where("address_longitude = ?", query.AddressLongitude)
	}
	if query.Moderated != nil {
		dbQuery = dbQuery.Where("moderated = ?", query.Moderated)
	}
	if query.Verified != nil {
		dbQuery = dbQuery.Where("verified = ?", query.Verified)
	}

	// Apply filtering for BodyArt and ProfileTags if present
	if len(query.BodyArtIds) > 0 {
		dbQuery = dbQuery.Joins("JOIN profile_body_arts ON profiles.id = profile_body_arts.profile_id").
			Where("profile_body_arts.body_art_id IN ?", query.BodyArtIds)
	}

	if len(query.ProfileTagIds) > 0 {
		dbQuery = dbQuery.Joins("JOIN profile_options ON profiles.id = profile_options.profile_id").
			Where("profile_options.profile_tag_id IN ?", query.ProfileTagIds)
	}

	// Apply price range filters with cases for nil min/max values
	if query.PriceInHouseContactMin != nil && query.PriceInHouseContactMax != nil {
		dbQuery = dbQuery.Where("price_in_house_contact BETWEEN ? AND ?", query.PriceInHouseContactMin, query.PriceInHouseContactMax)
	} else if query.PriceInHouseContactMin != nil {
		dbQuery = dbQuery.Where("price_in_house_contact >= ?", query.PriceInHouseContactMin)
	} else if query.PriceInHouseContactMax != nil {
		dbQuery = dbQuery.Where("price_in_house_contact <= ?", query.PriceInHouseContactMax)
	}

	if query.PriceInHouseHourMin != nil && query.PriceInHouseHourMax != nil {
		dbQuery = dbQuery.Where("price_in_house_hour BETWEEN ? AND ?", query.PriceInHouseHourMin, query.PriceInHouseHourMax)
	} else if query.PriceInHouseHourMin != nil {
		dbQuery = dbQuery.Where("price_in_house_hour >= ?", query.PriceInHouseHourMin)
	} else if query.PriceInHouseHourMax != nil {
		dbQuery = dbQuery.Where("price_in_house_hour <= ?", query.PriceInHouseHourMax)
	}

	if query.PriceSaunaContactMin != nil && query.PriceSaunaContactMax != nil {
		dbQuery = dbQuery.Where("price_sauna_contact BETWEEN ? AND ?", query.PriceSaunaContactMin, query.PriceSaunaContactMax)
	} else if query.PriceSaunaContactMin != nil {
		dbQuery = dbQuery.Where("price_sauna_contact >= ?", query.PriceSaunaContactMin)
	} else if query.PriceSaunaContactMax != nil {
		dbQuery = dbQuery.Where("price_sauna_contact <= ?", query.PriceSaunaContactMax)
	}

	if query.PriceSaunaHourMin != nil && query.PriceSaunaHourMax != nil {
		dbQuery = dbQuery.Where("price_sauna_hour BETWEEN ? AND ?", query.PriceSaunaHourMin, query.PriceSaunaHourMax)
	} else if query.PriceSaunaHourMin != nil {
		dbQuery = dbQuery.Where("price_sauna_hour >= ?", query.PriceSaunaHourMin)
	} else if query.PriceSaunaHourMax != nil {
		dbQuery = dbQuery.Where("price_sauna_hour <= ?", query.PriceSaunaHourMax)
	}

	if query.PriceVisitContactMin != nil && query.PriceVisitContactMax != nil {
		dbQuery = dbQuery.Where("price_visit_contact BETWEEN ? AND ?", query.PriceVisitContactMin, query.PriceVisitContactMax)
	} else if query.PriceVisitContactMin != nil {
		dbQuery = dbQuery.Where("price_visit_contact >= ?", query.PriceVisitContactMin)
	} else if query.PriceVisitContactMax != nil {
		dbQuery = dbQuery.Where("price_visit_contact <= ?", query.PriceVisitContactMax)
	}

	if query.PriceVisitHourMin != nil && query.PriceVisitHourMax != nil {
		dbQuery = dbQuery.Where("price_visit_hour BETWEEN ? AND ?", query.PriceVisitHourMin, query.PriceVisitHourMax)
	} else if query.PriceVisitHourMin != nil {
		dbQuery = dbQuery.Where("price_visit_hour >= ?", query.PriceVisitHourMin)
	} else if query.PriceVisitHourMax != nil {
		dbQuery = dbQuery.Where("price_visit_hour <= ?", query.PriceVisitHourMax)
	}

	if query.PriceCarContactMin != nil && query.PriceCarContactMax != nil {
		dbQuery = dbQuery.Where("price_car_contact BETWEEN ? AND ?", query.PriceCarContactMin, query.PriceCarContactMax)
	} else if query.PriceCarContactMin != nil {
		dbQuery = dbQuery.Where("price_car_contact >= ?", query.PriceCarContactMin)
	} else if query.PriceCarContactMax != nil {
		dbQuery = dbQuery.Where("price_car_contact <= ?", query.PriceCarContactMax)
	}

	if query.PriceCarHourMin != nil && query.PriceCarHourMax != nil {
		dbQuery = dbQuery.Where("price_car_hour BETWEEN ? AND ?", query.PriceCarHourMin, query.PriceCarHourMax)
	} else if query.PriceCarHourMin != nil {
		dbQuery = dbQuery.Where("price_car_hour >= ?", query.PriceCarHourMin)
	} else if query.PriceCarHourMax != nil {
		dbQuery = dbQuery.Where("price_car_hour <= ?", query.PriceCarHourMax)
	}

	currentUser := ctx.MustGet("currentUser").(User)
	if currentUser.Role == "user" {
		dbQuery = dbQuery.Where("active = ?", true)
	}

	// Execute the query
	results := dbQuery.Find(&profiles)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{
			Status:  "error",
			Message: results.Error.Error(),
		})
		return
	}

	intPage, _ = strconv.Atoi(page)

	profileResponses := make([]ProfileResponse, len(profiles))
	for i, profile := range profiles {
		profileResponses[i] = *utils.MapProfile(&profile) // Assuming you have the mapProfile function
	}

	// Return the results in the response
	ctx.JSON(http.StatusOK, SuccessPageResponse[[]ProfileResponse]{
		Status:  "success",
		Results: len(profiles),
		Page:    intPage,
		Data:    profileResponses,
	})
}

// DeleteProfile godoc
//
//	@Summary		Deletes a profile by ID
//	@Description	Deletes the profile with the given ID from the database
//	@Tags			Profiles
//	@Produce		json
//	@Param			id	path		string	true	"Profile ID"
//	@Success		204	{object}	nil
//	@Failure		404	{object}	ErrorResponse
//	@Router			/profiles/{id} [delete]
func (pc *ProfileController) DeleteProfile(ctx *gin.Context) {
	profileId := ctx.Param("id")

	result := pc.DB.Delete(&Profile{}, "id = ?", profileId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Status: "error", Message: "No profile with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
