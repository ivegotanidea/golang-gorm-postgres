package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ProfileOption struct {
	ProfileID    uuid.UUID      `gorm:"primaryKey;type:uuid"`    // Part of composite primary key
	ProfileTagID int            `gorm:"primaryKey;type:integer"` // Part of composite primary key, and also a foreign key
	Price        int64          `gorm:"type:bigint"`
	Comment      string         `gorm:"type:text"`
	ProfileTag   ProfileTag     `gorm:"foreignKey:ProfileTagID;references:ID"`
	Flags        datatypes.JSON `gorm:"index;type:jsonb;default:'{}'::jsonb"` // JSON column
}

type CreateProfileOptionFlagRequest struct {
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
	Price   *int64 `json:"price"`
}

type CreateProfileOptionRequest struct {
	ProfileTagID int                              `json:"profileTagId" binding:"required" validate:"required,int,gte=0"`
	Price        int64                            `json:"price,omitempty" validate:"min=0"`
	Comment      string                           `json:"comment,omitempty" validate:"max=50"`
	Flags        []CreateProfileOptionFlagRequest `json:"flags"`
}

type ProfileOptionFlagResponse struct {
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
	Price   *int64 `json:"price"`
}

type ProfileOptionResponse struct {
	Price   *int64                      `json:"price"`
	Comment string                      `json:"comment"`
	Name    string                      `json:"name"`
	Flags   []ProfileOptionFlagResponse `json:"flags"`
}
