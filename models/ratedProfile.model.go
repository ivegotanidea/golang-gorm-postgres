package models

import (
	"github.com/google/uuid"
)

type RatedProfileTag struct {
	RatingID     uuid.UUID  `gorm:"primaryKey;type:uuid"`
	ProfileTagID uuid.UUID  `gorm:"primaryKey;type:uuid"`
	Type         string     `gorm:"type:varchar(10)"`
	ProfileTag   ProfileTag `gorm:"foreignKey:ProfileTagID"`
}
