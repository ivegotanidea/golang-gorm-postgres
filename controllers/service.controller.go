package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type ServiceController struct {
	DB                     *gorm.DB
	reviewUpdateLimitHours int
}

func NewServiceController(DB *gorm.DB, reviewUpdateLimitHours int) ServiceController {
	return ServiceController{DB, reviewUpdateLimitHours}
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func getDistanceBetweenCoordinates(latA, lonA, latB, lonB float32) float64 {
	const earthRadiusKm = 6371

	dLat := degToRad(float64(latB - latA))
	dLon := degToRad(float64(lonB - lonA))

	latARad := degToRad(float64(latA))
	latBRad := degToRad(float64(latB))

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(latARad)*math.Cos(latBRad)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadiusKm * c

	return distance
}

func (sc *ServiceController) CreateService(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateServiceRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()

	distance := getDistanceBetweenCoordinates(
		*payload.ClientUserLatitude, *payload.ClientUserLongitude,
		*payload.ProfileUserLatitude, *payload.ProfileUserLongitude)

	// Create the new service object
	newService := models.Service{
		ClientUserID:  payload.ClientUserID,
		ClientUserLat: strconv.FormatFloat(float64(*payload.ClientUserLatitude), 'f', -1, 32),
		ClientUserLon: strconv.FormatFloat(float64(*payload.ClientUserLongitude), 'f', -1, 32),

		ProfileID:      payload.ProfileID,
		ProfileOwnerID: payload.ProfileOwnerID,
		ProfileUserLat: strconv.FormatFloat(float64(*payload.ProfileUserLatitude), 'f', -1, 32),
		ProfileUserLon: strconv.FormatFloat(float64(*payload.ProfileUserLongitude), 'f', -1, 32),

		DistanceBetweenUsers: distance,
		TrustedDistance:      distance <= 100,

		CreatedAt: now,
		UpdatedAt: now,
		UpdatedBy: currentUser.ID,
	}

	tx := sc.DB.Begin()

	// Create the service entry first
	err := tx.Create(&newService).Error
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create service"})
		return
	}

	// If profile rating exists
	if payload.ProfileRating != nil {
		// Create the profile rating and capture its ID
		reviewOfProfile := models.ProfileRating{
			ServiceID: newService.ID, // link to the service
			ProfileID: newService.ProfileID,
			Review:    payload.ProfileRating.Review,
			Score:     payload.ProfileRating.Score,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := tx.Create(&reviewOfProfile).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create profile rating"})
			return
		}

		// Capture the ProfileRating ID and assign it to the service
		newService.ProfileRatingID = &reviewOfProfile.ID

		// Handle RatedProfileTags creation
		if len(payload.ProfileRating.RatedProfileTags) > 0 {
			var ratedProfileTags []models.RatedProfileTag

			for _, profileTag := range payload.ProfileRating.RatedProfileTags {
				profileTagID := profileTag.TagID

				ratedProfileTag := models.RatedProfileTag{
					RatingID:     reviewOfProfile.ID, // reference the rating ID
					ProfileTagID: profileTagID,
					Type:         profileTag.Type,
				}
				ratedProfileTags = append(ratedProfileTags, ratedProfileTag)
			}

			if err := tx.Create(&ratedProfileTags).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create rated profile tags"})
				return
			}
		}
	}

	// If user rating exists
	if payload.UserRating != nil {
		// Create the user rating and capture its ID
		reviewOfUser := models.UserRating{
			ServiceID: newService.ID, // link to the service
			UserID:    newService.ClientUserID,
			Review:    payload.UserRating.Review,
			Score:     payload.UserRating.Score,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := tx.Create(&reviewOfUser).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create user rating"})
			return
		}

		// Capture the UserRating ID and assign it to the service
		newService.ClientUserRatingID = &reviewOfUser.ID

		// Handle RatedUserTags creation
		if len(payload.UserRating.RatedUserTags) > 0 {
			var ratedUserTags []models.RatedUserTag

			for _, userTag := range payload.UserRating.RatedUserTags {
				ratedUserTag := models.RatedUserTag{
					RatingID:  reviewOfUser.ID, // reference the rating ID
					UserTagID: userTag.TagID,
					Type:      userTag.Type,
				}
				ratedUserTags = append(ratedUserTags, ratedUserTag)
			}

			if err := tx.Create(&ratedUserTags).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create rated user tags"})
				return
			}
		}
	}

	// Commit the transaction
	if err := tx.Save(&newService).Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to commit transaction"})
		return
	}

	// Return the created service in the response
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newService})
}

func MutateService(tier string, service models.Service) map[string]interface{} {

	filteredService := make(map[string]interface{})

	// Common fields that are visible to all
	filteredService["ID"] = service.ID
	filteredService["ClientUserID"] = service.ClientUserID
	filteredService["ClientUserRatingID"] = service.ClientUserRatingID
	filteredService["ProfileID"] = service.ProfileID
	filteredService["ProfileOwnerID"] = service.ProfileOwnerID

	if service.ProfileRating != nil {
		filteredService["ProfileRatingID"] = service.ProfileRating.ID
	}

	filteredService["CreatedAt"] = service.CreatedAt
	filteredService["DistanceBetweenUsers"] = service.DistanceBetweenUsers
	filteredService["TrustedDistance"] = service.TrustedDistance

	// Access control based on user tier
	if tier == "basic" {
		// Only expose ProfileRating.Score for basic users, hide UserRating entirely
		if service.ProfileRating != nil {
			filteredService["ProfileRating"] = map[string]interface{}{
				"Score":             service.ProfileRating.Score,
				"ReviewTextVisible": true,
			}
		}
	} else if tier == "expert" {
		// Expert users can see all ProfileRating fields and only Score from UserRating
		if service.ProfileRating != nil {
			filteredService["ProfileRating"] = map[string]interface{}{
				"Review": service.ProfileRating.Review,
				"Score":  service.ProfileRating.Score,
			}
		}
		if service.ClientUserRating != nil {
			filteredService["ClientUserRating"] = map[string]interface{}{
				"Score":             service.ClientUserRating.Score,
				"ReviewTextVisible": true,
			}
		}
	} else if tier == "guru" {
		// Guru users can see both ProfileRating and UserRating completely
		filteredService["ProfileRating"] = service.ProfileRating
		filteredService["ClientUserRating"] = service.ClientUserRating
	}

	return filteredService
}

func (sc *ServiceController) GetService(ctx *gin.Context) {
	profileID := ctx.Param("profileID")
	serviceID := ctx.Param("serviceID")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var service models.Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("profile_id = ? and id = ?", profileID, serviceID).
		First(&service)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for specified profile"})
		return
	}

	filteredService := MutateService(currentUser.Tier, service)

	// Return the filtered response
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": filteredService})
}

func (sc *ServiceController) GetProfileServices(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	profileID := ctx.Param("profileID")
	currentUser := ctx.MustGet("currentUser").(models.User)

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var services []models.Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("profile_id = ?", profileID).
		Limit(intLimit).Offset(offset).
		Find(&services)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for specified profile"})
		return
	}

	filteredServices := make([]map[string]interface{}, 0, len(services))

	for _, service := range services {
		filteredService := MutateService(currentUser.Tier, service)
		// Append filtered service to the response
		filteredServices = append(filteredServices, filteredService)
	}

	// Return the filtered response
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(filteredServices), "data": filteredServices})
}

func (sc *ServiceController) ListServices(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []models.Service

	dbQuery := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&profiles)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(profiles), "data": profiles})
}

// ----

func (sc *ServiceController) UpdateClientUserReviewOnProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	serviceID := ctx.Query("service_id")

	// Find the service with the associated user review
	var service models.Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("id = ?", serviceID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for the specified profile"})
		return
	}

	// Check if the current user is the one who left the review
	if service.ClientUserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this review"})
		return
	}

	// Check if the review can still be updated (within the allowed time limit)
	hoursSinceReview := time.Since(service.ClientUserRating.CreatedAt).Hours()
	if hoursSinceReview > float64(sc.reviewUpdateLimitHours) {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": fmt.Sprintf("Review can only be updated within %d hours of creation", sc.reviewUpdateLimitHours)})
		return
	}

	// Parse the request payload
	var payload *models.CreateUserRatingRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Update the review fields
	if payload.Review != "" {
		service.ClientUserRating.Review = payload.Review
	}
	if payload.Score != nil {
		service.ClientUserRating.Score = payload.Score
	}

	// Handle the rated user tags if they exist in the payload
	if len(payload.RatedUserTags) > 0 {
		// First, delete the existing tags for this user rating
		if err := sc.DB.Where("rating_id = ?", service.ClientUserRating.ID).Delete(&models.RatedUserTag{}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old user tags"})
			return
		}

		// Iterate over the new tags and add them
		var ratedUserTags []models.RatedUserTag
		for _, tagReq := range payload.RatedUserTags {
			ratedUserTag := models.RatedUserTag{
				RatingID:  service.ClientUserRating.ID,
				UserTagID: tagReq.TagID,
				Type:      tagReq.Type,
			}
			ratedUserTags = append(ratedUserTags, ratedUserTag)
		}

		// Save the new tags
		if err := sc.DB.Create(&ratedUserTags).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create new user tags"})
			return
		}

		// Assign the new tags to the rating
		service.ClientUserRating.RatedUserTags = ratedUserTags
	}

	// Update the user rating in the database
	if err := sc.DB.Save(&service.ClientUserRating).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update the user review"})
		return
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": service})
}

func (sc *ServiceController) HideProfileOwnerReview(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	profileID := ctx.Param("profileID")

	// Find the service with the associated user review
	var service models.Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("profile_id = ?", profileID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for the specified profile"})
		return
	}

	// Check if the current user is the one who left the review
	if service.ClientUserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this review"})
		return
	}

	if !service.ClientUserRating.ReviewTextHidden {
		service.ClientUserRating.ReviewTextHidden = true

		if err := sc.DB.Save(&service.ClientUserRating).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update the user review"})
			return
		}
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": service})
}

func (sc *ServiceController) UpdateProfileOwnerReviewOnClientUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	serviceID := ctx.Query("service_id")

	// Find the service with the associated profile review
	var service models.Service
	result := sc.DB.Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Preload("ClientUserRating.RatedUserTags.UserTag").
		Where("id = ?", serviceID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for the specified profile"})
		return
	}

	// Check if the current user is the owner of the profile in the service
	if service.ProfileOwnerID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this review"})
		return
	}

	// Check if the review can still be updated (within the allowed time limit)
	hoursSinceReview := time.Since(service.ProfileRating.CreatedAt).Hours()

	if hoursSinceReview > float64(sc.reviewUpdateLimitHours) {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": fmt.Sprintf("Review can only be updated within %d hours of creation", sc.reviewUpdateLimitHours)})
		return
	}

	// Parse the request payload
	var payload *models.CreateProfileRatingRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Update the profile review fields
	if payload.Review != "" {
		service.ProfileRating.Review = payload.Review
	}
	if payload.Score != nil {
		service.ProfileRating.Score = payload.Score
	}

	// Handle the rated profile tags if they exist in the payload
	if len(payload.RatedProfileTags) > 0 {
		// First, delete the existing tags for this profile rating
		if err := sc.DB.Where("rating_id = ?", service.ProfileRating.ID).Delete(&models.RatedProfileTag{}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete old profile tags"})
			return
		}

		// Iterate over the new tags and add them
		var ratedProfileTags []models.RatedProfileTag
		for _, tagReq := range payload.RatedProfileTags {
			ratedProfileTag := models.RatedProfileTag{
				RatingID:     service.ProfileRating.ID,
				ProfileTagID: tagReq.TagID,
				Type:         tagReq.Type,
			}
			ratedProfileTags = append(ratedProfileTags, ratedProfileTag)
		}

		// Save the new tags
		if err := sc.DB.Create(&ratedProfileTags).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create new profile tags"})
			return
		}

		// Assign the new tags to the profile rating
		service.ProfileRating.RatedProfileTags = ratedProfileTags
	}

	// Update the profile rating in the database
	if err := sc.DB.Save(&service.ProfileRating).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update the profile review"})
		return
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": service})
}

func (sc *ServiceController) HideUserReview(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	serviceID := ctx.Query("service_id")

	// Find the service with the associated profile review
	var service models.Service
	result := sc.DB.Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Preload("ClientUserRating.RatedUserTags.UserTag").
		Where("id = ?", serviceID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for the specified profile"})
		return
	}

	// Check if the current user is the owner of the profile in the service
	if service.ProfileOwnerID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this review"})
		return
	}

	if currentUser.Tier == "basic" {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Basic-tier user can't hide out profile reviews"})
	}

	var payload *models.SetReviewVisibilityRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if service.ProfileRating.ReviewTextVisible != *payload.Visible {
		service.ProfileRating.ReviewTextVisible = *payload.Visible
		service.ProfileRating.UpdatedAt = time.Now()
		service.UpdatedBy = currentUser.ID

		// Update the profile rating in the database
		if err := sc.DB.Save(&service.ProfileRating).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update the profile review"})
			return
		}
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": service})
}
