package models

import (
	"github.com/google/uuid"
)

type ProfileOption struct {
	ProfileID    uuid.UUID  `gorm:"type:uuid;primaryKey"` // Part of composite primary key
	ProfileTagID uuid.UUID  `gorm:"type:uuid;primaryKey"` // Part of composite primary key, and also a foreign key
	Price        int64      `gorm:"type:bigint"`
	Comment      string     `gorm:"type:text"`
	ProfileTag   ProfileTag `gorm:"foreignKey:ProfileTagID;references:ID"`
}

type CreateProfileOption struct {
	ProfileID    uuid.UUID `json:"profileId" binding:"required" validate:"required,uuid"`
	ProfileTagID uuid.UUID `json:"profileTagId" binding:"required" validate:"required,uuid"`
	Price        int64     `json:"price,omitempty" validate:"min=0"`
	Comment      string    `json:"comment,omitempty" validate:"max=50"`
}
