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

type BodyType struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
}

type BodyArt struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:100;not null;unique"`
	AliasRu string `gorm:"size:100;not null"`
	AliasEn string `gorm:"size:100;not null"`
}

type ProfileBodyArt struct {
	ProfileID uuid.UUID `gorm:"primaryKey;type:uuid"` // Part of composite primary key
	BodyArtID int       `gorm:"primaryKey"`           // Part of composite primary key, and also a foreign key
}

type CreateBodyArtRequest struct {
	ID int `json:"bodyArtId" binding:"required"`
}

type HairColor struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
}

type IntimateHairCut struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
}

type ProfileBodyArtResponse struct {
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}
