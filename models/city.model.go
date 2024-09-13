package models

type City struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:30;not null;unique"`
	AliasRu string `gorm:"size:30;not null"`
	AliasEn string `gorm:"size:30;not null"`
}
