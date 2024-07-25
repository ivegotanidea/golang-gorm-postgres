package models

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	ID                 uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt          time.Time     `gorm:"type:timestamp"`
	ClientUserID       uuid.UUID     `gorm:"type:uuid;not null"`
	ClientUserRatingID uuid.UUID     `gorm:"type:uuid"`
	ClientUserRating   UserRating    `gorm:"foreignKey:ClientUserRatingID"`
	ClientUserLat      string        `gorm:"type:varchar(10)"`
	ClientUserLon      string        `gorm:"type:varchar(10)"`
	ProfileID          uuid.UUID     `gorm:"type:uuid;not null"`
	ProfileRatingID    uuid.UUID     `gorm:"type:uuid"`
	ProfileRating      ProfileRating `gorm:"foreignKey:ProfileRatingID"`
	ProfileUserLat     string        `gorm:"type:varchar(10)"`
	ProfileUserLon     string        `gorm:"type:varchar(10)"`
}
