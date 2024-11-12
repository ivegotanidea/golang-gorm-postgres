package models

import (
	"github.com/google/uuid"
)

type ProfileOption struct {
	ProfileID    uuid.UUID  `gorm:"primaryKey;type:uuid"`    // Part of composite primary key
	ProfileTagID int        `gorm:"primaryKey;type:integer"` // Part of composite primary key, and also a foreign key
	Price        int64      `gorm:"type:bigint"`
	Comment      string     `gorm:"type:text"`
	ProfileTag   ProfileTag `gorm:"foreignKey:ProfileTagID;references:ID"`
}

type CreateProfileOptionRequest struct {
	ProfileTagID int    `json:"profileTagId" binding:"required" validate:"required,int,gte=0"`
	Price        int64  `json:"price,omitempty" validate:"min=0"`
	Comment      string `json:"comment,omitempty" validate:"max=50"`
}

type ProfileOptionResponse struct {
	Price      int64              `json:"price"`
	Comment    string             `json:"comment"`
	ProfileTag ProfileTagResponse `json:"profileTag"`
}
