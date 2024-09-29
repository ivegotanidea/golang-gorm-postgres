package controllers

import (
	"fmt"
	"github.com/google/uuid"
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

	err := tx.Create(&newService).Error

	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create service"})
		return
	}

	if payload.ProfileRating != nil {
		reviewOfProfile := models.ProfileRating{
			ServiceID: newService.ID,
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

		// Handle RatedProfileTags creation
		if len(payload.ProfileRating.RatedProfileTags) > 0 {
			var ratedProfileTags []models.RatedProfileTag

			for _, profileTag := range payload.ProfileRating.RatedProfileTags {
				profileTagID, _ := uuid.Parse(profileTag.TagID)

				ratedProfileTag := models.RatedProfileTag{
					RatingID:     reviewOfProfile.ID,
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

		newService.ProfileRatingID = &reviewOfProfile.ID
	}

	if payload.UserRating != nil {
		reviewOfUser := models.UserRating{
			ServiceID: newService.ID,
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

		// Handle RatedUserTags creation
		if len(payload.UserRating.RatedUserTags) > 0 {
			var ratedUserTags []models.RatedUserTag

			for _, userTag := range payload.UserRating.RatedUserTags {
				userTagID, _ := uuid.Parse(userTag.TagID)

				ratedUserTag := models.RatedUserTag{
					RatingID:  reviewOfUser.ID,
					UserTagID: userTagID,
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

		newService.ClientUserRatingID = &reviewOfUser.ID
	}

	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to commit transaction"})
		return
	}

	// Return the created service in the response
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newService})
}

func (sc *ServiceController) UpdateClientUserReviewOnProfile(ctx *gin.Context) {
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
			tagID, _ := uuid.Parse(tagReq.TagID)
			ratedUserTag := models.RatedUserTag{
				RatingID:  service.ClientUserRating.ID,
				UserTagID: tagID,
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
	profileID := ctx.Param("profileID")

	// Find the service with the associated profile review
	var service models.Service
	result := sc.DB.Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Preload("ClientUserRating.RatedUserTags.UserTag").
		Where("profile_id = ?", profileID).
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
			tagID, _ := uuid.Parse(tagReq.TagID)
			ratedProfileTag := models.RatedProfileTag{
				RatingID:     service.ProfileRating.ID,
				ProfileTagID: tagID,
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
	profileID := ctx.Param("profileID")

	// Find the service with the associated profile review
	var service models.Service
	result := sc.DB.Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Preload("ClientUserRating.RatedUserTags.UserTag").
		Where("profile_id = ?", profileID).
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

	if !service.ProfileRating.ReviewTextHidden {
		service.ProfileRating.ReviewTextHidden = true

		// Update the profile rating in the database
		if err := sc.DB.Save(&service.ProfileRating).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update the profile review"})
			return
		}
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": service})
}

func (sc *ServiceController) GetProfileServices(ctx *gin.Context) {
	profileID := ctx.Param("profileID")

	var services []models.Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("profile_id = ?", profileID).
		Find(&services)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for specified profile"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(services), "data": services})
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
