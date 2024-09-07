package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	BodyTypeId        uint `gorm:"type:integer;default:null"`
	EthnosId          uint `gorm:"type:integer;default:null"` // todo: not null by default
	HairColorId       uint `gorm:"type:integer;default:null"`
	IntimateHairCutId uint `gorm:"type:integer;default:null"`

	// deprecate
	Ethnos string `gorm:"type:varchar(30);not null"`

	CityID           uint      `gorm:"type:integer; not null"`
	ParsedUrl        string    `gorm:"type:varchar(255);not null;default:''"`
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID           uuid.UUID `gorm:"type:uuid;not null"`
	Active           bool      `gorm:"type:boolean;default:true"`
	Phone            string    `gorm:"type:varchar(30)"` // ;index:,unique,composite:idx_single_profile"
	Name             string    `gorm:"type:varchar(50)"`
	Age              int       `gorm:"type:int;not null"`
	Height           int       `gorm:"type:int;not null"`
	Weight           int       `gorm:"type:int;not null"`
	Bust             float64   `gorm:"type:float"`
	Bio              string    `gorm:"type:varchar(2000)"`
	AddressLatitude  string    `gorm:"type:varchar(10)"`
	AddressLongitude string    `gorm:"type:varchar(10)"`

	PriceInHouseNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceInHouseContact    int     `gorm:"type:int"`
	PriceInHouseHour       int     `gorm:"type:int"`

	PrinceSaunaNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceSaunaContact     int     `gorm:"type:int"`
	PriceSaunaHour        int     `gorm:"type:int"`

	PriceVisitNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceVisitContact    int     `gorm:"type:int"`
	PriceVisitHour       int     `gorm:"type:int"`

	PriceCarNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceCarContact    int     `gorm:"type:int"`
	PriceCarHour       int     `gorm:"type:int"`

	ContactPhone string `gorm:"type:varchar(30)"`
	ContactWA    string `gorm:"type:varchar(30)"`
	ContactTG    string `gorm:"type:varchar(50)"`

	Moderated bool `gorm:"type:boolean;default:false"`
	Verified  bool `gorm:"type:boolean;default:false"`

	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedBy uuid.UUID      `gorm:"type:uuid;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	BodyArt        []ProfileBodyArt `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE;"`
	Photos         []Photo          `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE;"`
	ProfileOptions []ProfileOption  `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE;"`
	Services       []Service        `gorm:"foreignKey:ProfileID"`
}

type CreateProfileRequest struct {
	Phone  string  `json:"phone"  binding:"required"`
	Name   string  `json:"name"  binding:"required"`
	Age    int     `json:"age"  binding:"required"`
	CityID uint    `json:"cityId"  binding:"required"`
	Height int     `json:"height"  binding:"required"`
	Weight int     `json:"weight"  binding:"required"`
	Bust   float64 `json:"bust"  binding:"required"`
	Ethnos string  `json:"ethnos"  binding:"omitempty"`

	EthnosId          uint `json:"ethnosId"  binding:"required"`
	HairColorId       uint `json:"hairColorId"  binding:"omitempty"`
	BodyTypeId        uint `json:"bodyTypeId"  binding:"omitempty"`
	IntimateHairCutId uint `json:"intimateHairCutId"  binding:"omitempty"`

	Bio string `json:"bio"  binding:"required"`

	AddressLatitude  string `json:"latitude,omitempty"`
	AddressLongitude string `json:"longitude,omitempty"`

	//PriceInHouseNightRatio float64 `json:"priceInHouseNightRatio,omitempty"`
	PriceInHouseContact int `json:"priceInHouseContact,omitempty"`
	PriceInHouseHour    int `json:"priceInHouseHour,omitempty"`
	//PrinceSaunaNightRatio  float64 `json:"princeSaunaNightRatio,omitempty"`
	PriceSaunaContact int `json:"priceSaunaContact,omitempty"`
	PriceSaunaHour    int `json:"priceSaunaHour,omitempty"`
	//PriceVisitNightRatio   float64 `json:"priceVisitNightRatio,omitempty"`
	PriceVisitContact int `json:"priceVisitContact,omitempty"`
	PriceVisitHour    int `json:"priceVisitHour,omitempty"`
	//PriceCarNightRatio     float64 `json:"priceCarNightRatio,omitempty"`
	PriceCarContact int `json:"priceCarContact,omitempty"`
	PriceCarHour    int `json:"priceCarHour,omitempty"`

	ContactPhone string `json:"contactPhone" binding:"required"`
	ContactTG    string `json:"contactTG" binding:"required"`
	ContactWA    string `json:"contactWA,omitempty"`

	BodyArts []CreateBodyArtRequest `json:"bodyArts" binding:"omitempty"`
	Photos   []CreatePhotoRequest   `json:"photos" binding:"required,dive"`
	Options  []CreateProfileOption  `json:"profileOptions" binding:"required,dive"`
}

type UpdateProfileRequest struct {
	CityID     int     `json:"cityId"  binding:"required"`
	Phone      string  `json:"phone"  binding:"required"`
	Name       string  `json:"name"  binding:"required"`
	Age        int     `json:"age"  binding:"required"`
	Height     int     `json:"height"  binding:"required"`
	Weight     int     `json:"weight"  binding:"required"`
	Bust       float64 `json:"bust"  binding:"required"`
	BodyTypeId int     `json:"bodyTypeId"  binding:"omitempty"`
	Ethnos     string  `json:"ethnos"  binding:"required"`
	Bio        string  `json:"bio"  binding:"required"`

	PriceInHouseNightRatio float64 `json:"priceInHouseNightRatio,omitempty"`
	PriceInHouseContact    int     `json:"priceInHouseContact,omitempty"`
	PriceInHouseHour       int     `json:"priceInHouseHour,omitempty"`
	PrinceSaunaNightRatio  float64 `json:"princeSaunaNightRatio,omitempty"`
	PriceSaunaContact      int     `json:"priceSaunaContact,omitempty"`
	PriceSaunaHour         int     `json:"priceSaunaHour,omitempty"`
	PriceVisitNightRatio   float64 `json:"priceVisitNightRatio,omitempty"`
	PriceVisitContact      int     `json:"priceVisitContact,omitempty"`
	PriceVisitHour         int     `json:"priceVisitHour,omitempty"`
	PriceCarNightRatio     float64 `json:"priceCarNightRatio,omitempty"`
	PriceCarContact        int     `json:"priceCarContact,omitempty"`
	PriceCarHour           int     `json:"priceCarHour,omitempty"`

	ContactPhone string `json:"contactPhone" binding:"required"`
	ContactTG    string `json:"contactTG" binding:"required"`
	ContactWA    string `json:"contactWA,omitempty"`

	Photos  []CreatePhotoRequest  `json:"photos" binding:"required,dive"`
	Options []CreateProfileOption `json:"profileOptions" binding:"required,dive"`
}
