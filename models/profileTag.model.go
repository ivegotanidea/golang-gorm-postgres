package models

import "gorm.io/datatypes"

type ProfileTag struct {
	ID      int            `gorm:"type:integer;primaryKey"`
	Name    string         `gorm:"size:100;not null;unique"`
	AliasRu string         `gorm:"size:100;"`
	AliasEn string         `gorm:"size:100;"`
	Flags   datatypes.JSON `gorm:"index;type:jsonb;default:'{}'::jsonb"` // JSON column
}

type ProfileTagFlagResponse struct {
	Name    string `json:"name"`
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
