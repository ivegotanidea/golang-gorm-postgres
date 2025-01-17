package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"time"
)

type ProfileTag struct {
	ID      int            `gorm:"type:integer;primaryKey"`
	Name    string         `gorm:"size:100;not null;unique"`
	AliasRu string         `gorm:"size:100;"`
	AliasEn string         `gorm:"size:100;"`
	Flags   datatypes.JSON `gorm:"index;type:jsonb;default:'{}'::jsonb"` // JSON column

	// todo set not null
	CreatedAt time.Time `gorm:"type:timestamp;"`
	CreatedBy uuid.UUID `json:"createdBy"`

	UpdatedAt time.Time `gorm:"type:timestamp;"`
	UpdatedBy uuid.UUID `json:"updatedBy"`
}

type ProfileTagFlagResponse struct {
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}

type ProfileTagResponse struct {
	ID      int                      `json:"id"`
	Name    string                   `json:"name"`
	AliasRu string                   `json:"aliasRu"`
	AliasEn string                   `json:"aliasEn"`
	Flags   []ProfileTagFlagResponse `json:"flags"`
}
