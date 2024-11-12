package models

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	ClientUserID       uuid.UUID   `gorm:"type:uuid;not null"`
	ClientUserRatingID *uuid.UUID  `gorm:"type:uuid"`
	ClientUserRating   *UserRating `gorm:"foreignKey:ClientUserRatingID"` // profile owner's review on client
	ClientUserLat      string      `gorm:"type:varchar(10)"`
	ClientUserLon      string      `gorm:"type:varchar(10)"`

	ProfileID       uuid.UUID      `gorm:"type:uuid;not null"`
	ProfileOwnerID  uuid.UUID      `gorm:"type:uuid;not null"`
	ProfileRatingID *uuid.UUID     `gorm:"type:uuid"`
	ProfileRating   *ProfileRating `gorm:"foreignKey:ProfileRatingID"` // client's review on profile
	ProfileUserLat  string         `gorm:"type:varchar(10)"`
	ProfileUserLon  string         `gorm:"type:varchar(10)"`

	DistanceBetweenUsers float64
	TrustedDistance      bool

	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null"`
}

type CreateServiceRequest struct {
	ClientUserID        uuid.UUID `json:"userId" validate:"gte=0"`
	ClientUserLatitude  *float32  `json:"clientUserLatitude" validate:"latitude"`
	ClientUserLongitude *float32  `json:"clientUserLongitude" validate:"longitude"`

	ProfileID            uuid.UUID `json:"profileId" validate:"gte=0"`
	ProfileOwnerID       uuid.UUID `json:"profileOwnerId" validate:"gte=0"`
	ProfileUserLatitude  *float32  `json:"profileUserLatitude" validate:"latitude"`
	ProfileUserLongitude *float32  `json:"profileUserLongitude" validate:"longitude"`

	UserRating    *CreateUserRatingRequest    `json:"userRating" binding:"omitempty"`
	ProfileRating *CreateProfileRatingRequest `json:"profileRating" binding:"omitempty"`
}

type CreateRatedUserTagRequest struct {
	Type  string `json:"type" validate:"required,oneOf=like dislike"`
	TagID int    `json:"tagId" validate:"required,gte=0"`
}

type CreateUserRatingRequest struct {
	Review        string                      `json:"review" binding:"omitempty" validate:"max=2000"`
	Score         *int                        `json:"score" binding:"omitempty" validate:"min=0,max=5"`
	RatedUserTags []CreateRatedUserTagRequest `json:"ratedUserTags" binding:"omitempty,dive"`
}

type CreateRatedProfileTagRequest struct {
	Type  string `json:"type" validate:"required,oneOf=like dislike"`
	TagID int    `json:"tagId" validate:"required,gte=0"`
}

type CreateProfileRatingRequest struct {
	Review           string                         `json:"review" binding:"omitempty" validate:"max=2000"`
	Score            *int                           `json:"score" binding:"omitempty" validate:"min=0,max=5"`
	RatedProfileTags []CreateRatedProfileTagRequest `json:"ratedProfileTags" binding:"omitempty,dive"`
}

type SetReviewVisibilityRequest struct {
	Visible *bool `json:"visible" validate:"boolean"`
}

type ServiceResponse struct {
	ID                   uuid.UUID              `json:"id"`
	ClientUserID         uuid.UUID              `json:"clientUserId"`
	ClientUserRatingID   *uuid.UUID             `json:"clientUserRatingId,omitempty"`
	ClientUserRating     *UserRatingResponse    `json:"clientUserRating,omitempty"`
	ProfileID            uuid.UUID              `json:"profileId"`
	ProfileOwnerID       uuid.UUID              `json:"profileOwnerId"`
	ProfileRatingID      *uuid.UUID             `json:"profileRatingId,omitempty"`
	ProfileRating        *ProfileRatingResponse `json:"profileRating,omitempty"`
	DistanceBetweenUsers float64                `json:"distanceBetweenUsers"`
	TrustedDistance      bool                   `json:"trustedDistance"`
	CreatedAt            time.Time              `json:"createdAt"`
	UpdatedAt            time.Time              `json:"updatedAt"`
	UpdatedBy            uuid.UUID              `json:"updatedBy"`
}
