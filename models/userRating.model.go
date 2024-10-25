package models

import (
	"github.com/google/uuid"
	"time"
)

type UserRating struct {
	ID                uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ServiceID         uuid.UUID      `gorm:"type:uuid;not null"`
	UserID            uuid.UUID      `gorm:"type:uuid;not null"`
	ReviewTextVisible bool           `gorm:"default:true"`
	Review            string         `gorm:"type:varchar(2000)"`
	Score             *int           `gorm:"type:int;not null"`
	CreatedAt         time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt         time.Time      `gorm:"type:timestamp;not null"`
	RatedUserTags     []RatedUserTag `gorm:"foreignKey:RatingID"`

	UpdatedBy uuid.UUID `gorm:"type:uuid;not null"`
}

type UserRatingResponse struct {
	ID                uuid.UUID              `json:"id"`
	ServiceID         uuid.UUID              `json:"serviceId"`
	UserID            uuid.UUID              `json:"userId"`
	ReviewTextVisible bool                   `json:"reviewTextVisible"`
	Review            string                 `json:"review"`
	Score             *int                   `json:"score"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         time.Time              `json:"updatedAt"`
	RatedUserTags     []RatedUserTagResponse `json:"ratedUserTags"`
	UpdatedBy         uuid.UUID              `json:"updatedBy"`
}
