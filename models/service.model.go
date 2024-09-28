package models

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID                 uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ClientUserID       uuid.UUID      `gorm:"type:uuid;not null"`
	ClientUserRatingID *uuid.UUID     `gorm:"type:uuid"`
	ClientUserRating   *UserRating    `gorm:"foreignKey:ClientUserRatingID"`
	ClientUserLat      string         `gorm:"type:varchar(10)"`
	ClientUserLon      string         `gorm:"type:varchar(10)"`
	ProfileID          uuid.UUID      `gorm:"type:uuid;not null"`
	ProfileRatingID    *uuid.UUID     `gorm:"type:uuid"`
	ProfileRating      *ProfileRating `gorm:"foreignKey:ProfileRatingID"`
	ProfileUserLat     string         `gorm:"type:varchar(10)"`
	ProfileUserLon     string         `gorm:"type:varchar(10)"`

	DistanceBetweenUsers float64
	TrustedDistance      bool

	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null"`
}

type CreateServiceRequest struct {
	ClientUserID        uuid.UUID `json:"userId"`
	ClientUserLatitude  *float32  `json:"clientUserLatitude"`
	ClientUserLongitude *float32  `json:"clientUserLongitude"`

	ProfileID            uuid.UUID `json:"profileId"`
	ProfileUserLatitude  *float32  `json:"profileUserLatitude"`
	ProfileUserLongitude *float32  `json:"profileUserLongitude"`

	UserRating    *CreateUserRatingRequest    `json:"userRating" binding:"omitempty,dive"`
	ProfileRating *CreateProfileRatingRequest `json:"profileRating" binding:"omitempty,dive"`
}

type CreateRatedUserTagRequest struct {
	Type  string `json:"type"`
	TagID string `json:"tagId"`
}

type CreateUserRatingRequest struct {
	Review        string                      `json:"review" binding:"omitempty"`
	Score         *int                        `json:"score" binding:"omitempty"`
	RatedUserTags []CreateRatedUserTagRequest `json:"ratedUserTags" binding:"omitempty,dive"`
}

type CreateRatedProfileTagRequest struct {
	Type  string `json:"type"`
	TagID string `json:"tagId"`
}

type CreateProfileRatingRequest struct {
	Review           string                         `json:"review" binding:"omitempty"`
	Score            *int                           `json:"score" binding:"omitempty"`
	RatedProfileTags []CreateRatedProfileTagRequest `json:"ratedProfileTags" binding:"omitempty,dive"`
}
