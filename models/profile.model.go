package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	BodyTypeID        *int `gorm:"type:integer;default:null"`
	EthnosID          *int `gorm:"type:integer;default:null"`
	HairColorID       *int `gorm:"type:integer;default:null"`
	IntimateHairCutID *int `gorm:"type:integer;default:null"`

	CityID           int       `gorm:"type:integer;not null;default:0"`
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
	PriceInHouseContact    *int    `gorm:"type:int;default:null"`
	PriceInHouseHour       *int    `gorm:"type:int;default:null"`

	PrinceSaunaNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceSaunaContact     *int    `gorm:"type:int;default:null"`
	PriceSaunaHour        *int    `gorm:"type:int;default:null"`

	PriceVisitNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceVisitContact    *int    `gorm:"type:int;default:null"`
	PriceVisitHour       *int    `gorm:"type:int;default:null"`

	PriceCarNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceCarContact    *int    `gorm:"type:int;default:null"`
	PriceCarHour       *int    `gorm:"type:int;default:null"`

	ContactPhone string `gorm:"type:varchar(30)"`
	ContactWA    string `gorm:"type:varchar(30)"`
	ContactTG    string `gorm:"type:varchar(50)"`

	Moderated   bool      `gorm:"type:boolean;default:false"`
	ModeratedAt time.Time `gorm:"type:timestamp;default:null"`
	ModeratedBy uuid.UUID `gorm:"type:uuid;default:null"`

	Verified   bool      `gorm:"type:boolean;default:false"`
	VerifiedAt time.Time `gorm:"type:timestamp;default:null"`
	VerifiedBy uuid.UUID `gorm:"type:uuid;default:null"`

	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedBy uuid.UUID      `gorm:"type:uuid;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`

	BodyArts       []ProfileBodyArt `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE;"`
	Photos         []Photo          `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE;"`
	ProfileOptions []ProfileOption  `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE;"`
	Services       []Service        `gorm:"foreignKey:ProfileID"`
}

type CreateProfileRequest struct {
	Phone  string  `json:"phone"  binding:"required"`
	Name   string  `json:"name"  binding:"required"`
	Age    int     `json:"age"  binding:"required"`
	CityID int     `json:"cityId"  binding:"required"`
	Height int     `json:"height"  binding:"required"`
	Weight int     `json:"weight"  binding:"required"`
	Bust   float64 `json:"bust"  binding:"required"`

	EthnosID          *int `json:"ethnosId"  binding:"required"`
	HairColorID       *int `json:"hairColorId"  binding:"omitempty"`
	BodyTypeID        *int `json:"bodyTypeId"  binding:"omitempty"`
	IntimateHairCutID *int `json:"intimateHairCutId"  binding:"omitempty"`

	Bio string `json:"bio"  binding:"required"`

	AddressLatitude  string `json:"latitude,omitempty"`
	AddressLongitude string `json:"longitude,omitempty"`

	//PriceInHouseNightRatio float64 `json:"priceInHouseNightRatio,omitempty"`
	PriceInHouseContact *int `json:"priceInHouseContact,omitempty"`
	PriceInHouseHour    *int `json:"priceInHouseHour,omitempty"`
	//PrinceSaunaNightRatio  float64 `json:"princeSaunaNightRatio,omitempty"`
	PriceSaunaContact *int `json:"priceSaunaContact,omitempty"`
	PriceSaunaHour    *int `json:"priceSaunaHour,omitempty"`
	//PriceVisitNightRatio   float64 `json:"priceVisitNightRatio,omitempty"`
	PriceVisitContact *int `json:"priceVisitContact,omitempty"`
	PriceVisitHour    *int `json:"priceVisitHour,omitempty"`
	//PriceCarNightRatio     float64 `json:"priceCarNightRatio,omitempty"`
	PriceCarContact *int `json:"priceCarContact,omitempty"`
	PriceCarHour    *int `json:"priceCarHour,omitempty"`

	ContactPhone string `json:"contactPhone" binding:"required"`
	ContactTG    string `json:"contactTG" binding:"required"`
	ContactWA    string `json:"contactWA,omitempty"`

	BodyArts []CreateBodyArtRequest `json:"bodyArts" binding:"omitempty,dive"`
	Photos   []CreatePhotoRequest   `json:"photos" binding:"required,dive"`
	Options  []CreateProfileOption  `json:"profileOptions" binding:"required,dive"`
}

type UpdateOwnProfileRequest struct {
	Active *bool    `json:"active" binding:"omitempty"`
	CityID *int     `json:"cityId"  binding:"omitempty"`
	Phone  string   `json:"phone"  binding:"omitempty"`
	Name   string   `json:"name"  binding:"omitempty"`
	Age    *int     `json:"age"  binding:"omitempty"`
	Height *int     `json:"height"  binding:"omitempty"`
	Weight *int     `json:"weight"  binding:"omitempty"`
	Bust   *float64 `json:"bust"  binding:"omitempty"`

	BodyTypeID        *int `json:"bodyTypeId"  binding:"omitempty"`
	EthnosID          *int `json:"ethnosId"  binding:"omitempty"`
	HairColorID       *int `json:"hairColorId,omitempty"`
	IntimateHairCutID *int `json:"intimateHairCutId,omitempty"`

	Bio string `json:"bio"  binding:"omitempty"`

	AddressLatitude  string `json:"latitude,omitempty"`
	AddressLongitude string `json:"longitude,omitempty"`

	PriceInHouseNightRatio *float64 `json:"priceInHouseNightRatio,omitempty"`
	PriceInHouseContact    *int     `json:"priceInHouseContact,omitempty"`
	PriceInHouseHour       *int     `json:"priceInHouseHour,omitempty"`
	PrinceSaunaNightRatio  *float64 `json:"princeSaunaNightRatio,omitempty"`
	PriceSaunaContact      *int     `json:"priceSaunaContact,omitempty"`
	PriceSaunaHour         *int     `json:"priceSaunaHour,omitempty"`
	PriceVisitNightRatio   *float64 `json:"priceVisitNightRatio,omitempty"`
	PriceVisitContact      *int     `json:"priceVisitContact,omitempty"`
	PriceVisitHour         *int     `json:"priceVisitHour,omitempty"`
	PriceCarNightRatio     *float64 `json:"priceCarNightRatio,omitempty"`
	PriceCarContact        *int     `json:"priceCarContact,omitempty"`
	PriceCarHour           *int     `json:"priceCarHour,omitempty"`

	ContactPhone string `json:"contactPhone" binding:"omitempty"`
	ContactTG    string `json:"contactTG" binding:"omitempty"`
	ContactWA    string `json:"contactWA,omitempty"`

	BodyArts []CreateBodyArtRequest `json:"bodyArts" binding:"omitempty,dive"`
	Photos   []CreatePhotoRequest   `json:"photos" binding:"omitempty,dive"`
	Options  []CreateProfileOption  `json:"profileOptions" binding:"omitempty,dive"`
}

type UpdateProfileRequest struct {
	Active    *bool  `json:"active" binding:"omitempty"`
	Name      string `json:"name"  binding:"omitempty"`
	Bio       string `json:"bio"  binding:"omitempty"`
	Moderated *bool  `json:"moderated" binding:"omitempty"`
	Verified  *bool  `json:"verified" binding:"omitempty"`

	Photos []CreatePhotoRequest `json:"photos" binding:"omitempty,dive"`
}

type FindProfilesQuery struct {
	BodyTypeId             *int     `json:"bodyTypeId,omitempty"`
	EthnosId               *int     `json:"ethnosId,omitempty"`
	HairColorId            *int     `json:"hairColorId,omitempty"`
	IntimateHairCutId      *int     `json:"intimateHairCutId,omitempty"`
	CityID                 *int     `json:"cityId,omitempty"`
	Active                 *bool    `json:"active,omitempty"`
	Phone                  string   `json:"phone,omitempty"`
	Age                    *int     `json:"age,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Height                 *int     `json:"height,omitempty"`
	Weight                 *int     `json:"weight,omitempty"`
	Bust                   *float64 `json:"bust,omitempty"`
	AddressLatitude        string   `json:"latitude,omitempty"`
	AddressLongitude       string   `json:"longitude,omitempty"`
	Moderated              *bool    `json:"moderated,omitempty"`
	Verified               *bool    `json:"verified,omitempty"`
	BodyArtIds             []*int   `json:"bodyArtIds,omitempty"`
	ProfileTagIds          []*int   `json:"profileTagIds,omitempty"`
	PriceInHouseContactMin *int     `json:"priceInHouseContactMin,omitempty"`
	PriceInHouseContactMax *int     `json:"priceInHouseContactMax,omitempty"`
	PriceInHouseHourMin    *int     `json:"priceInHouseHourMin,omitempty"`
	PriceInHouseHourMax    *int     `json:"priceInHouseHourMax,omitempty"`
	PriceSaunaContactMin   *int     `json:"priceSaunaContactMin,omitempty"`
	PriceSaunaContactMax   *int     `json:"priceSaunaContactMax,omitempty"`
	PriceSaunaHourMin      *int     `json:"priceSaunaHourMin,omitempty"`
	PriceSaunaHourMax      *int     `json:"priceSaunaHourMax,omitempty"`
	PriceVisitContactMin   *int     `json:"priceVisitContactMin,omitempty"`
	PriceVisitContactMax   *int     `json:"priceVisitContactMax,omitempty"`
	PriceVisitHourMin      *int     `json:"priceVisitHourMin,omitempty"`
	PriceVisitHourMax      *int     `json:"priceVisitHourMax,omitempty"`
	PriceCarContactMin     *int     `json:"priceCarContactMin,omitempty"`
	PriceCarContactMax     *int     `json:"priceCarContactMax,omitempty"`
	PriceCarHourMin        *int     `json:"priceCarHourMin,omitempty"`
	PriceCarHourMax        *int     `json:"priceCarHourMax,omitempty"`
}

type ProfileResponse struct {
	UserID                 string                   `json:"userId"`
	Active                 bool                     `json:"active"`
	Phone                  string                   `json:"phone"`
	Name                   string                   `json:"name"`
	Age                    int                      `json:"age"`
	Height                 int                      `json:"height"`
	Weight                 int                      `json:"weight"`
	Bust                   float64                  `json:"bust"`
	Bio                    string                   `json:"bio"`
	AddressLatitude        string                   `json:"addressLatitude,omitempty"`
	AddressLongitude       string                   `json:"addressLongitude,omitempty"`
	CityID                 *int                     `json:"cityId,omitempty"`
	BodyTypeID             *int                     `json:"bodyTypeId,omitempty"`
	EthnosID               *int                     `json:"ethnosId,omitempty"`
	HairColorID            *int                     `json:"hairColorId,omitempty"`
	IntimateHairCutID      *int                     `json:"intimateHairCutId,omitempty"`
	PriceInHouseNightRatio float64                  `json:"priceInHouseNightRatio"`
	PriceInHouseContact    *int                     `json:"priceInHouseContact,omitempty"`
	PriceInHouseHour       *int                     `json:"priceInHouseHour,omitempty"`
	PrinceSaunaNightRatio  float64                  `json:"princeSaunaNightRatio"`
	PriceSaunaContact      *int                     `json:"priceSaunaContact,omitempty"`
	PriceSaunaHour         *int                     `json:"priceSaunaHour,omitempty"`
	PriceVisitNightRatio   float64                  `json:"priceVisitNightRatio"`
	PriceVisitContact      *int                     `json:"priceVisitContact,omitempty"`
	PriceVisitHour         *int                     `json:"priceVisitHour,omitempty"`
	PriceCarNightRatio     float64                  `json:"priceCarNightRatio"`
	PriceCarContact        *int                     `json:"priceCarContact,omitempty"`
	PriceCarHour           *int                     `json:"priceCarHour,omitempty"`
	ContactPhone           string                   `json:"contactPhone"`
	ContactWA              string                   `json:"contactWA,omitempty"`
	ContactTG              string                   `json:"contactTG,omitempty"`
	Moderated              bool                     `json:"moderated"`
	ModeratedAt            *time.Time               `json:"moderatedAt,omitempty"`
	Verified               bool                     `json:"verified"`
	VerifiedAt             *time.Time               `json:"verifiedAt,omitempty"`
	VerifiedBy             *uuid.UUID               `json:"verifiedBy,omitempty"`
	CreatedAt              time.Time                `json:"createdAt"`
	BodyArts               []ProfileBodyArtResponse `json:"bodyArts,omitempty"`
	Photos                 []PhotoResponse          `json:"photos,omitempty"`
	ProfileOptions         []ProfileOptionResponse  `json:"profileOptions,omitempty"`
	Services               []ServiceResponse        `json:"services,omitempty"`
	UpdatedBy              *uuid.UUID               `json:"updatedBy,omitempty"`
}
