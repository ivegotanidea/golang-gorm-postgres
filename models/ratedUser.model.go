package models

import (
	"github.com/google/uuid"
)

type RatedUserTag struct {
	RatingID  uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserTagID uuid.UUID `gorm:"primaryKey;type:uuid"`
	Type      string    `gorm:"type:varchar(10)"`
	UserTag   UserTag   `gorm:"foreignKey:UserTagID"`
}
