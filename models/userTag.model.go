package models

type UserTag struct {
	ID   int    `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

type UserTagResponse struct {
	Name string `json:"name"`
}
