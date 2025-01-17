package models

import (
	"github.com/google/uuid"
	"time"
)

type UserTag struct {
	ID      int    `gorm:"type:integer;primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;"`
	AliasEn string `gorm:"size:30;"`

	// todo set not null
	CreatedAt time.Time `gorm:"type:timestamp;"`
	CreatedBy uuid.UUID `json:"createdBy"`

	UpdatedAt time.Time `gorm:"type:timestamp;"`
	UpdatedBy uuid.UUID `json:"updatedBy"`
}

type UserTagResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}
