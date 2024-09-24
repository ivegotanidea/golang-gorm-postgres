package controllers

import (
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
	DB *gorm.DB
}

func NewServiceController(DB *gorm.DB) ServiceController {
	return ServiceController{DB}
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

func (sc *ServiceController) UpdateUserReview(ctx *gin.Context) {

}

func (sc *ServiceController) UpdateProfileOwnerReview(ctx *gin.Context) {

}

func (sc *ServiceController) GetProfileServices(ctx *gin.Context) {
	profileID := ctx.Param("profileID")

	var services []models.Service
	result := sc.DB.Preload("ClientUserRating.RatedUserTags.UserTag").
		Preload("ProfileRating.RatedProfileTags.ProfileTag").
		Where("profile_id = ?", profileID).
		Find(&services)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No services found for the specified profile"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(services), "data": services})
}

func (sc *ServiceController) UpdateService(ctx *gin.Context) {
	postId := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePost
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedPost models.Post
	result := sc.DB.First(&updatedPost, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: now,
	}

	sc.DB.Model(&updatedPost).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})
}

func (sc *ServiceController) ListServices(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []models.Service

	dbQuery := sc.DB.Preload("Photos").
		Preload("BodyArts").
		Preload("ProfileOptions.ProfileTag").
		Limit(intLimit).Offset(offset)

	results := dbQuery.Find(&profiles)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(profiles), "data": profiles})
}
