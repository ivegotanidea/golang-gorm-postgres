package models

type ProfileTag struct {
	ID      int    `gorm:"type:integer;primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;"`
	AliasEn string `gorm:"size:30;"`
}

type ProfileTagResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AliasRu string `json:"aliasRu"`
	AliasEn string `json:"aliasEn"`
}
