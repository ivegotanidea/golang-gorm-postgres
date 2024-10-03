package models

import (
	"github.com/google/uuid"
	"time"
)

type UserRating struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ServiceID        uuid.UUID      `gorm:"type:uuid;not null"`
	UserID           uuid.UUID      `gorm:"type:uuid;not null"`
	ReviewTextHidden bool           `gorm:"default:false"`
	Review           string         `gorm:"type:varchar(2000)"`
	Score            *int           `gorm:"type:int;not null"`
	CreatedAt        time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt        time.Time      `gorm:"type:timestamp;not null"`
	RatedUserTags    []RatedUserTag `gorm:"foreignKey:RatingID"`

	UpdatedBy uuid.UUID `gorm:"type:uuid;not null"`
}
