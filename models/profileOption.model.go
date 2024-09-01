package models

import (
	"github.com/google/uuid"
)

type ProfileOption struct {
	ProfileID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	ProfileTagID uuid.UUID `gorm:"foreignKey:ProfileTagID"`
	Price        int64     `gorm:"type:bigint"`
	Comment      string    `gorm:"type:text"`
}

type CreateProfileOption struct {
	ProfileID    uuid.UUID `json:"profileId" binding:"required" validate:"required,uuid"`
	ProfileTagID uuid.UUID `json:"profileTagId" binding:"required" validate:"required,uuid"`
	Price        int64     `json:"price,omitempty" validate:"min=0"`
	Comment      string    `json:"comment,omitempty" validate:"max=50"`
}
