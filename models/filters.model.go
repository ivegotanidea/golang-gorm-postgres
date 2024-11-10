package models

import (
	"github.com/google/uuid"
)

type Ethnos struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
	Sex     string `gorm:"size:10;not null"`
}

type EthnosResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
	Sex     string `json:"sex"`
}

type BodyType struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
}

type BodyTypeResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}

type BodyArt struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:100;not null;unique"`
	AliasRu string `gorm:"size:100;not null"`
	AliasEn string `gorm:"size:100;not null"`
}

type BodyArtResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}

type ProfileBodyArt struct {
	ProfileID uuid.UUID `gorm:"primaryKey;type:uuid"` // Part of composite primary key
	BodyArtID int       `gorm:"primaryKey"`           // Part of composite primary key, and also a foreign key
}

type CreateBodyArtRequest struct {
	ID int `json:"bodyArtId" binding:"required" validate:"gte=0"`
}

type HairColor struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
}

// HairColorResponse represents the JSON response for hair color data
type HairColorResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}

type IntimateHairCut struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
}

// IntimateHairCutResponse represents the JSON response for intimate hair cut data
type IntimateHairCutResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}

type ProfileBodyArtResponse struct {
	ProfileID string `json:"profileId"`
	BodyArtID int    `json:"bodyArtId"`
}
