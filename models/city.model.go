package models

import (
	"github.com/google/uuid"
	"time"
)

type City struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`

	// todo set not null
	CreatedAt time.Time `gorm:"type:timestamp;"`
	CreatedBy uuid.UUID `json:"createdBy"`

	UpdatedAt time.Time `gorm:"type:timestamp;"`
	UpdatedBy uuid.UUID `json:"updatedBy"`
}

type CityResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}
