package models

import (
	"github.com/google/uuid"
	"time"
)

type ProfileRating struct {
	ID                uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ServiceID         uuid.UUID         `gorm:"type:uuid;not null"`
	ProfileID         uuid.UUID         `gorm:"type:uuid;not null"`
	ReviewTextVisible bool              `gorm:"default:true"`
	Review            string            `gorm:"type:varchar(2000)"`
	Score             *int              `gorm:"type:int;not null"`
	CreatedAt         time.Time         `gorm:"type:timestamp;not null"`
	UpdatedAt         time.Time         `gorm:"type:timestamp;not null"`
	RatedProfileTags  []RatedProfileTag `gorm:"foreignKey:RatingID"`

	UpdatedBy uuid.UUID `gorm:"type:uuid;not null"`
}

type ProfileRatingResponse struct {
	ID                uuid.UUID                 `json:"id"`
	ServiceID         uuid.UUID                 `json:"serviceId"`
	ProfileID         uuid.UUID                 `json:"profileId"`
	ReviewTextVisible bool                      `json:"reviewTextVisible"`
	Review            string                    `json:"review"`
	Score             *int                      `json:"score"`
	CreatedAt         time.Time                 `json:"createdAt"`
	UpdatedAt         time.Time                 `json:"updatedAt"`
	RatedProfileTags  []RatedProfileTagResponse `json:"ratedProfileTags"`
	UpdatedBy         uuid.UUID                 `json:"updatedBy"`
}
