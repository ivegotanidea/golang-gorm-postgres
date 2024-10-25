package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
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

// CreateService godoc
//
//	@Summary		Create a new service
//	@Description	Creates a new service between a client user and a profile, including optional ratings for both the profile and the user.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateServiceRequest	true	"Create Service Request"
//	@Success		201		{object}	SuccessResponse{data=Service}
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/services [post]
func (sc *ServiceController) CreateService(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(User)
	var payload *CreateServiceRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	now := time.Now()

	distance := getDistanceBetweenCoordinates(
		*payload.ClientUserLatitude, *payload.ClientUserLongitude,
		*payload.ProfileUserLatitude, *payload.ProfileUserLongitude)

	// Create the new service object
	newService := Service{
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

	// Create the service entry
	err := tx.Create(&newService).Error
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to create service",
		})
		return
	}

	// If profile rating exists
	if payload.ProfileRating != nil {
		reviewOfProfile := ProfileRating{
			ServiceID: newService.ID,
			ProfileID: newService.ProfileID,
			Review:    payload.ProfileRating.Review,
			Score:     payload.ProfileRating.Score,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := tx.Create(&reviewOfProfile).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to create profile rating",
			})
			return
		}

		newService.ProfileRatingID = &reviewOfProfile.ID

		if len(payload.ProfileRating.RatedProfileTags) > 0 {
			var ratedProfileTags []RatedProfileTag
			for _, profileTag := range payload.ProfileRating.RatedProfileTags {
				ratedProfileTags = append(ratedProfileTags, RatedProfileTag{
					RatingID:     reviewOfProfile.ID,
					ProfileTagID: profileTag.TagID,
					Type:         profileTag.Type,
				})
			}

			if err := tx.Create(&ratedProfileTags).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{
					Status:  "error",
					Message: "Failed to create rated profile tags",
				})
				return
			}
		}
	}

	// If user rating exists
	if payload.UserRating != nil {
		reviewOfUser := UserRating{
			ServiceID: newService.ID,
			UserID:    newService.ClientUserID,
			Review:    payload.UserRating.Review,
			Score:     payload.UserRating.Score,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := tx.Create(&reviewOfUser).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to create user rating",
			})
			return
		}

		newService.ClientUserRatingID = &reviewOfUser.ID

		if len(payload.UserRating.RatedUserTags) > 0 {
			var ratedUserTags []RatedUserTag
			for _, userTag := range payload.UserRating.RatedUserTags {
				ratedUserTags = append(ratedUserTags, RatedUserTag{
					RatingID:  reviewOfUser.ID,
					UserTagID: userTag.TagID,
					Type:      userTag.Type,
				})
			}

			if err := tx.Create(&ratedUserTags).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{
					Status:  "error",
					Message: "Failed to create rated user tags",
				})
				return
			}
		}
	}

	// Commit the transaction
	if err := tx.Save(&newService).Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to commit transaction",
		})
		return
	}

	// Return the created service in the response
	ctx.JSON(http.StatusCreated, SuccessResponse{
		Status: "success",
		Data:   newService,
	})
}

func MutateService(tier string, service Service) map[string]interface{} {

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

// GetService godoc
//
//	@Summary		Get a specific service by profile and service ID
//	@Description	Retrieves a service based on the profile ID and service ID, with filtered data based on the user's tier.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			profileID	path		string	true	"Profile ID"
//	@Param			serviceID	path		string	true	"Service ID"
//	@Success		200			{object}	SuccessResponse{data=Service}
//	@Failure		404			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/profiles/{profileID}/services/{serviceID} [get]
func (sc *ServiceController) GetService(ctx *gin.Context) {
	profileID := ctx.Param("profileID")
	serviceID := ctx.Param("serviceID")
	currentUser := ctx.MustGet("currentUser").(User)

	var service Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("profile_id = ? and id = ?", profileID, serviceID).
		First(&service)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Status:  "error",
			Message: "No services found for specified profile",
		})
		return
	}

	// Mutate the service based on the user's tier
	filteredService := MutateService(currentUser.Tier, service)

	// Return the filtered response
	ctx.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   filteredService,
	})
}

// GetProfileServices godoc
//
//	@Summary		Get all services for a specific profile
//	@Description	Retrieves all services for a specific profile, with filtered data based on the user's tier.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			profileID	path		string	true	"Profile ID"
//	@Param			page		query		string	false	"Page number"				default(1)
//	@Param			limit		query		string	false	"Number of items per page"	default(10)
//	@Success		200			{object}	SuccessResponse{data=[]map[string]interface{}}
//	@Failure		404			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/profiles/{profileID}/services [get]
func (sc *ServiceController) GetProfileServices(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	profileID := ctx.Param("profileID")
	currentUser := ctx.MustGet("currentUser").(User)

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var services []Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("profile_id = ?", profileID).
		Limit(intLimit).Offset(offset).
		Find(&services)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Status:  "error",
			Message: "No services found for specified profile",
		})
		return
	}

	// Mutate services based on user tier
	filteredServices := make([]map[string]interface{}, 0, len(services))
	for _, service := range services {
		filteredService := MutateService(currentUser.Tier, service)
		filteredServices = append(filteredServices, filteredService)
	}

	// Return the filtered response
	ctx.JSON(http.StatusOK, SuccessPageResponse{
		Status:  "success",
		Results: len(filteredServices),
		Page:    intPage,
		Limit:   intLimit,
		Data:    filteredServices,
	})
}

// ListServices godoc
//
//	@Summary		Get a list of services
//	@Description	Retrieves a paginated list of services with all related information.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"Page number"				default(1)
//	@Param			limit	query		string	false	"Number of items per page"	default(10)
//	@Success		200		{object}	SuccessResponse{data=[]Service}
//	@Failure		502		{object}	ErrorResponse
//	@Router			/services [get]
func (sc *ServiceController) ListServices(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var services []Service

	dbQuery := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&services)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponse{
			Status:  "error",
			Message: results.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, SuccessPageResponse{
		Status:  "success",
		Results: len(services),
		Limit:   intLimit,
		Page:    intPage,
		Data:    services,
	})
}

// ----

// UpdateClientUserReviewOnProfile godoc
//
//	@Summary		Update a client's user review on a profile service
//	@Description	Updates the user review for a service if the current user is authorized to do so and within the allowed time limit.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			serviceId	query		string					true	"Service ID"
//	@Param			body		body		CreateUserRatingRequest	true	"User Rating Request"
//	@Success		200			{object}	SuccessResponse{data=Service}
//	@Failure		400			{object}	ErrorResponse
//	@Failure		403			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/reviews/client/update [put]
func (sc *ServiceController) UpdateClientUserReviewOnProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(User)
	serviceID := ctx.Query("serviceId")

	// Find the service with the associated user review
	var service Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("id = ?", serviceID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Status:  "error",
			Message: "No services found for the specified profile",
		})
		return
	}

	// Check if the current user is the one who left the review
	if service.ClientUserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, ErrorResponse{
			Status:  "error",
			Message: "You are not authorized to update this review",
		})
		return
	}

	// Check if the review can still be updated (within the allowed time limit)
	hoursSinceReview := time.Since(service.ClientUserRating.CreatedAt).Hours()

	if hoursSinceReview > float64(sc.reviewUpdateLimitHours) {
		ctx.JSON(http.StatusForbidden, ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Review can only be updated within %d hours of creation", sc.reviewUpdateLimitHours),
		})
		return
	}

	// Parse the request payload
	var payload *CreateUserRatingRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
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
		if err := sc.DB.Where("rating_id = ?", service.ClientUserRating.ID).Delete(&RatedUserTag{}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to delete old user tags",
			})
			return
		}

		// Iterate over the new tags and add them
		var ratedUserTags []RatedUserTag
		for _, tagReq := range payload.RatedUserTags {
			ratedUserTag := RatedUserTag{
				RatingID:  service.ClientUserRating.ID,
				UserTagID: tagReq.TagID,
				Type:      tagReq.Type,
			}
			ratedUserTags = append(ratedUserTags, ratedUserTag)
		}

		// Save the new tags
		if err := sc.DB.Create(&ratedUserTags).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to create new user tags",
			})
			return
		}

		// Assign the new tags to the rating
		service.ClientUserRating.RatedUserTags = ratedUserTags
	}

	// Update the user rating in the database
	if err := sc.DB.Save(&service.ClientUserRating).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to update the user review",
		})
		return
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   service,
	})
}

// HideProfileOwnerReview godoc
//
//	@Summary		Set visibility of the profile owner's review
//	@Description	Set visibility of the profile owner's review based on the client's request. Only available for non-basic tier users.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			serviceId	query		string						true	"Service ID"
//	@Param			body		body		SetReviewVisibilityRequest	true	"Set Review Visibility Request"
//	@Success		200			{object}	SuccessResponse{data=Service}
//	@Failure		400			{object}	ErrorResponse
//	@Failure		403			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/reviews/client/visibility [put]
func (sc *ServiceController) HideProfileOwnerReview(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(User)
	serviceID := ctx.Query("serviceId")

	// Find the service with the associated user review
	var service Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("id = ?", serviceID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Status:  "error",
			Message: "No services found for the specified profile",
		})
		return
	}

	// Check if the current user is the one who left the review
	if service.ClientUserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, ErrorResponse{
			Status:  "error",
			Message: "You are not authorized to update this review",
		})
		return
	}

	var payload *SetReviewVisibilityRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// Update review visibility if changed
	if service.ClientUserRating.ReviewTextVisible != *payload.Visible {
		service.ClientUserRating.ReviewTextVisible = *payload.Visible

		if err := sc.DB.Save(&service.ClientUserRating).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to update the user review",
			})
			return
		}
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   service,
	})
}

// UpdateProfileOwnerReviewOnClientUser godoc
//
//	@Summary		Update the profile owner's review on a client user
//	@Description	Allows a profile owner to update their review on a client user within the allowed time limit.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			serviceId	query		string						true	"Service ID"
//	@Param			body		body		CreateProfileRatingRequest	true	"Create Profile Rating Request"
//	@Success		200			{object}	SuccessResponse{data=Service}
//	@Failure		400			{object}	ErrorResponse
//	@Failure		403			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/reviews/host/update [put]
func (sc *ServiceController) UpdateProfileOwnerReviewOnClientUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(User)
	serviceID := ctx.Query("serviceId")

	// Find the service with the associated profile review
	var service Service
	result := sc.DB.Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Preload("ClientUserRating.RatedUserTags.UserTag").
		Where("id = ?", serviceID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Status:  "error",
			Message: "No services found for the specified profile",
		})
		return
	}

	// Check if the current user is the owner of the profile in the service
	if service.ProfileOwnerID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, ErrorResponse{
			Status:  "error",
			Message: "You are not authorized to update this review",
		})
		return
	}

	// Check if the review can still be updated (within the allowed time limit)
	hoursSinceReview := time.Since(service.ProfileRating.CreatedAt).Hours()
	if hoursSinceReview > float64(sc.reviewUpdateLimitHours) {
		ctx.JSON(http.StatusForbidden, ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Review can only be updated within %d hours of creation", sc.reviewUpdateLimitHours),
		})
		return
	}

	// Parse the request payload
	var payload *CreateProfileRatingRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
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
		if err := sc.DB.Where("rating_id = ?", service.ProfileRating.ID).Delete(&RatedProfileTag{}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to delete old profile tags",
			})
			return
		}

		// Iterate over the new tags and add them
		var ratedProfileTags []RatedProfileTag
		for _, tagReq := range payload.RatedProfileTags {
			ratedProfileTag := RatedProfileTag{
				RatingID:     service.ProfileRating.ID,
				ProfileTagID: tagReq.TagID,
				Type:         tagReq.Type,
			}
			ratedProfileTags = append(ratedProfileTags, ratedProfileTag)
		}

		// Save the new tags
		if err := sc.DB.Create(&ratedProfileTags).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to create new profile tags",
			})
			return
		}

		// Assign the new tags to the profile rating
		service.ProfileRating.RatedProfileTags = ratedProfileTags
	}

	// Update the profile rating in the database
	if err := sc.DB.Save(&service.ProfileRating).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to update the profile review",
		})
		return
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   service,
	})
}

// HideUserReview godoc
//
//	@Summary		Set visibility of user's review
//	@Description	Allows a profile owner to set user's review visibility. Only available for non-basic tier users.
//	@Tags			Services
//	@Accept			json
//	@Produce		json
//	@Param			serviceId	query		string						true	"Service ID"
//	@Param			body		body		SetReviewVisibilityRequest	true	"Set Review Visibility Request"
//	@Success		200			{object}	SuccessResponse{data=Service}
//	@Failure		400			{object}	ErrorResponse
//	@Failure		403			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse
//	@Failure		500			{object}	ErrorResponse
//	@Router			/reviews/host/visibility [put]
func (sc *ServiceController) HideUserReview(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(User)
	serviceID := ctx.Query("serviceId")

	// Find the service with the associated profile review
	var service Service
	result := sc.DB.Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Preload("ClientUserRating.RatedUserTags.UserTag").
		Where("id = ?", serviceID).
		First(&service)

	// Check if the service exists
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{
			Status:  "error",
			Message: "No services found for the specified profile",
		})
		return
	}

	// Check if the current user is the owner of the profile in the service
	if service.ProfileOwnerID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, ErrorResponse{
			Status:  "error",
			Message: "You are not authorized to update this review",
		})
		return
	}

	// Check if the user is allowed to hide the review
	if currentUser.Tier == "basic" {
		ctx.JSON(http.StatusForbidden, ErrorResponse{
			Status:  "error",
			Message: "Basic-tier users can't hide profile reviews",
		})
		return
	}

	// Parse the request payload
	var payload *SetReviewVisibilityRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// Update review visibility if changed
	if service.ProfileRating.ReviewTextVisible != *payload.Visible {
		service.ProfileRating.ReviewTextVisible = *payload.Visible
		service.ProfileRating.UpdatedAt = time.Now()
		service.UpdatedBy = currentUser.ID

		// Update the profile rating in the database
		if err := sc.DB.Save(&service.ProfileRating).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  "error",
				Message: "Failed to update the profile review",
			})
			return
		}
	}

	// Return the updated service with the updated review
	ctx.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   service,
	})
}
