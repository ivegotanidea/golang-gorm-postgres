package models

type ProfileTag struct {
	ID   uint   `gorm:"type:integer;primaryKey"`
	Name string `gorm:"uniqueIndex;type:varchar(50)"`
}
