package models

import (
	"github.com/google/uuid"
)

type Profile struct {
	ID                  uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID              uuid.UUID       `gorm:"type:uuid;not null"`
	Active              bool            `gorm:"type:boolean"`
	Phone               string          `gorm:"type:varchar(30)"`
	Name                string          `gorm:"type:varchar(30)"`
	Age                 int             `gorm:"type:int;not null"`
	Height              int32           `gorm:"type:int;not null"`
	Weight              int32           `gorm:"type:int;not null"`
	Length              int32           `gorm:"type:int"`
	Bust                float64         `gorm:"type:float"`
	Ethnos              string          `gorm:"type:varchar(30);not null"`
	Bio                 string          `gorm:"type:varchar(2000)"`
	Moderated           bool            `gorm:"type:boolean"`
	Verified            bool            `gorm:"type:boolean"`
	PriceInHouseContact int             `gorm:"type:int"`
	PriceInHouseHour    int             `gorm:"type:int"`
	PriceSaunaContact   int             `gorm:"type:int"`
	PriceSaunaHour      int             `gorm:"type:int"`
	PriceVisitContact   int             `gorm:"type:int"`
	PriceVisitHour      int             `gorm:"type:int"`
	PriceCarContact     int             `gorm:"type:int"`
	PriceCarHour        int             `gorm:"type:int"`
	ContactPhone        string          `gorm:"type:varchar(30)"`
	ContactWA           string          `gorm:"type:varchar(30)"`
	ContactTG           string          `gorm:"type:varchar(50)"`
	Photos              []Photo         `gorm:"foreignKey:ProfileID"`
	ProfileOptions      []ProfileOption `gorm:"foreignKey:ProfileID"`
	Services            []Service       `gorm:"foreignKey:ProfileID"`
}
