package models

type ProfileTag struct {
	ID   int    `gorm:"type:integer;primaryKey"`
	Name string `gorm:"uniqueIndex;type:varchar(50)"`
}

type ProfileTagResponse struct {
	Name string `json:"name"`
}
