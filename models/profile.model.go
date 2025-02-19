package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	BodyTypeID        *int             `gorm:"type:integer;default:null"`
	BodyType          *BodyType        `gorm:"foreignKey:BodyTypeID"`
	EthnosID          *int             `gorm:"type:integer;default:null"`
	Ethnos            *Ethnos          `gorm:"foreignKey:EthnosID"`
	HairColorID       *int             `gorm:"type:integer;default:null"`
	HairColor         *HairColor       `gorm:"foreignKey:HairColorID"`
	IntimateHairCutID *int             `gorm:"type:integer;default:null"`
	IntimateHairCut   *IntimateHairCut `gorm:"foreignKey:IntimateHairCutID"`
	CityID            int              `gorm:"type:integer;not null;default:0"`
	City              *City            `gorm:"foreignKey:CityID"`
	ParsedUrl         string           `gorm:"type:varchar(255);default:null"`
	ParsedID          *int             `gorm:"type:integer;default:null;index:idx_parsed_id,unique,where:parsed_id IS NOT NULL"`
	ID                uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID            uuid.UUID        `gorm:"type:uuid;not null"`
	Active            bool             `gorm:"type:boolean;default:true"`
	Phone             string           `gorm:"type:varchar(30)"` // ;index:,unique,composite:idx_single_profile"
	Name              string           `gorm:"type:varchar(50)"`
	Age               int              `gorm:"type:int;not null"`
	Height            int              `gorm:"type:int;not null"`
	Weight            int              `gorm:"type:int;not null"`
	Bust              float64          `gorm:"type:float"`
	Bio               string           `gorm:"type:varchar(2000)"`
	Sex               string           `gorm:"type:varchar(10)"`
	AddressLatitude   string           `gorm:"type:varchar(10)"`
	AddressLongitude  string           `gorm:"type:varchar(10)"`

	PriceInHouseNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceInHouseContact    *int    `gorm:"type:int;default:null"`
	PriceInHouseHour       *int    `gorm:"type:int;default:null"`

	PriceSaunaNightRatio float64 `gorm:"type:float;not null;default:1"`
	PriceSaunaContact    *int    `gorm:"type:int;default:null"`
	PriceSaunaHour       *int    `gorm:"type:int;default:null"`

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
	Phone  string  `json:"phone"  binding:"required" validate:"e164"`
	Name   string  `json:"name"  binding:"required" validate:"min=3,max=20"`
	Sex    string  `json:"sex" binding:"required" validate:"min=3,max=10"`
	Age    int     `json:"age"  binding:"required" validate:"gte=18,lte=80"`
	CityID int     `json:"cityId"  binding:"required" validate:"gte=0"`
	Height int     `json:"height"  binding:"required" validate:"gte=0,lte=300"`
	Weight int     `json:"weight"  binding:"required" validate:"gte=0,lte=150"`
	Bust   float64 `json:"bust"  binding:"required" validate:"gte=0,lte=10"`

	ParsedUrl string `json:"parsedUrl" binding:"omitempty" validate:"uri"`
	ParsedID  *int   `json:"parsedId" binding:"omitempty" validate:"gte=0,lte=9999999"`

	EthnosID          *int `json:"ethnosId"  binding:"required" validate:"gte=0"`
	HairColorID       *int `json:"hairColorId"  binding:"omitempty" validate:"gte=0"`
	BodyTypeID        *int `json:"bodyTypeId"  binding:"omitempty" validate:"gte=0"`
	IntimateHairCutID *int `json:"intimateHairCutId"  binding:"omitempty" validate:"gte=0"`

	Bio string `json:"bio"  binding:"omitempty" validate:"min=100,max=2000"`

	AddressLatitude  string `json:"latitude,omitempty" validate:"latitude"`
	AddressLongitude string `json:"longitude,omitempty" validate:"longitude"`

	//PriceInHouseNightRatio float64 `json:"priceInHouseNightRatio,omitempty"`
	PriceInHouseContact *int `json:"priceInHouseContact,omitempty" validate:"gte=0"`
	PriceInHouseHour    *int `json:"priceInHouseHour,omitempty" validate:"gte=0"`
	//PriceSaunaNightRatio  float64 `json:"priceSaunaNightRatio,omitempty"`
	PriceSaunaContact *int `json:"priceSaunaContact,omitempty" validate:"gte=0"`
	PriceSaunaHour    *int `json:"priceSaunaHour,omitempty" validate:"gte=0"`
	//PriceVisitNightRatio   float64 `json:"priceVisitNightRatio,omitempty"`
	PriceVisitContact *int `json:"priceVisitContact,omitempty" validate:"gte=0"`
	PriceVisitHour    *int `json:"priceVisitHour,omitempty" validate:"gte=0"`
	//PriceCarNightRatio     float64 `json:"priceCarNightRatio,omitempty"`
	PriceCarContact *int `json:"priceCarContact,omitempty" validate:"gte=0"`
	PriceCarHour    *int `json:"priceCarHour,omitempty" validate:"gte=0"`

	ContactPhone string `json:"contactPhone" binding:"required" validate:"e164"`
	ContactTG    string `json:"contactTG" binding:"required" validate:"min=4"`
	ContactWA    string `json:"contactWA,omitempty" validate:"e164"`

	BodyArts []CreateBodyArtRequest       `json:"bodyArts" binding:"omitempty,dive"`
	Photos   []CreatePhotoRequest         `json:"photos" binding:"omitempty,dive"`
	Options  []CreateProfileOptionRequest `json:"profileOptions" binding:"required,dive"`
}

type UpdateOwnProfileRequest struct {
	Active *bool    `json:"active" binding:"omitempty" validate:"boolean"`
	CityID *int     `json:"cityId"  binding:"omitempty" validate:"gte=0"`
	Phone  string   `json:"phone"  binding:"omitempty" validate:"e164"`
	Name   string   `json:"name"  binding:"omitempty" validate:"min=3,max=20"`
	Age    *int     `json:"age"  binding:"omitempty" validate:"gte=18,lte=80"`
	Height *int     `json:"height"  binding:"omitempty" validate:"gte=0,lte=300"`
	Weight *int     `json:"weight"  binding:"omitempty" validate:"gte=0,lte=150"`
	Bust   *float64 `json:"bust"  binding:"omitempty" validate:"gte=0,lte=10"`

	BodyTypeID        *int `json:"bodyTypeId"  binding:"omitempty" validate:"gte=0"`
	EthnosID          *int `json:"ethnosId"  binding:"omitempty" validate:"gte=0"`
	HairColorID       *int `json:"hairColorId,omitempty" validate:"gte=0"`
	IntimateHairCutID *int `json:"intimateHairCutId,omitempty" validate:"gte=0"`

	Bio string `json:"bio"  binding:"omitempty" validate:"min=100,max=2000"`

	AddressLatitude  string `json:"latitude,omitempty" validate:"latitude"`
	AddressLongitude string `json:"longitude,omitempty" validate:"longitude"`

	PriceInHouseNightRatio *float64 `json:"priceInHouseNightRatio,omitempty" validate:"gte=0"`
	PriceInHouseContact    *int     `json:"priceInHouseContact,omitempty" validate:"gte=0"`
	PriceInHouseHour       *int     `json:"priceInHouseHour,omitempty" validate:"gte=0"`
	PriceSaunaNightRatio   *float64 `json:"priceSaunaNightRatio,omitempty" validate:"gte=0"`
	PriceSaunaContact      *int     `json:"priceSaunaContact,omitempty" validate:"gte=0"`
	PriceSaunaHour         *int     `json:"priceSaunaHour,omitempty" validate:"gte=0"`
	PriceVisitNightRatio   *float64 `json:"priceVisitNightRatio,omitempty" validate:"gte=0"`
	PriceVisitContact      *int     `json:"priceVisitContact,omitempty" validate:"gte=0"`
	PriceVisitHour         *int     `json:"priceVisitHour,omitempty" validate:"gte=0"`
	PriceCarNightRatio     *float64 `json:"priceCarNightRatio,omitempty" validate:"gte=0"`
	PriceCarContact        *int     `json:"priceCarContact,omitempty" validate:"gte=0"`
	PriceCarHour           *int     `json:"priceCarHour,omitempty" validate:"gte=0"`

	ContactPhone string `json:"contactPhone" binding:"omitempty" validate:"e164"`
	ContactTG    string `json:"contactTG" binding:"omitempty" validate:"min=4"`
	ContactWA    string `json:"contactWA,omitempty" validate:"e164"`

	BodyArts []CreateBodyArtRequest       `json:"bodyArts" binding:"omitempty,dive"`
	Photos   []CreatePhotoRequest         `json:"photos" binding:"omitempty,dive"`
	Options  []CreateProfileOptionRequest `json:"profileOptions" binding:"omitempty,dive"`
}

type UpdateProfileRequest struct {
	Active    *bool  `json:"active" binding:"omitempty" validate:"boolean"`
	Name      string `json:"name"  binding:"omitempty" validate:"min=3,max=20"`
	Bio       string `json:"bio"  binding:"omitempty" validate:"min=100,max=2000"`
	Moderated *bool  `json:"moderated" binding:"omitempty" validate:"boolean"`
	Verified  *bool  `json:"verified" binding:"omitempty" validate:"boolean"`

	Photos []CreatePhotoRequest `json:"photos" binding:"omitempty,dive"`
}

type ListProfilesQuery struct {
	Page   int    `form:"page" validate:"gte=0"`
	Limit  int    `form:"limit" validate:"gte=0;lte=12"`
	CityID int    `form:"city" validate:"gte=0;lte=100"`
	Sex    string `form:"sex" validate:"oneof=female male"`
}

type FindProfilesQuery struct {
	BodyTypeId             *int     `json:"bodyTypeId,omitempty" validate:"gte=0"`
	EthnosId               *int     `json:"ethnosId,omitempty" validate:"gte=0"`
	HairColorId            *int     `json:"hairColorId,omitempty" validate:"gte=0"`
	IntimateHairCutId      *int     `json:"intimateHairCutId,omitempty" validate:"gte=0"`
	CityID                 *int     `json:"cityId,omitempty" validate:"gte=0"`
	Active                 *bool    `json:"active,omitempty" validate:"boolean"`
	Phone                  string   `json:"phone,omitempty" validate:"e164"`
	Age                    *int     `json:"age,omitempty" validate:"gte=18,lte=80"`
	Name                   string   `json:"name,omitempty" validate:"min=3,max=20"`
	Height                 *int     `json:"height,omitempty" validate:"gte=0,lte=300"`
	Weight                 *int     `json:"weight,omitempty" validate:"gte=0,lte=150"`
	Bust                   *float64 `json:"bust,omitempty" validate:"gte=0,lte=10"`
	AddressLatitude        string   `json:"latitude,omitempty" validate:"latitude"`
	AddressLongitude       string   `json:"longitude,omitempty" validate:"longitude"`
	Moderated              *bool    `json:"moderated,omitempty" validate:"boolean"`
	Verified               *bool    `json:"verified,omitempty" validate:"boolean"`
	BodyArtIds             []*int   `json:"bodyArtIds,omitempty" validate:"dive gte=0"`
	ProfileTagIds          []*int   `json:"profileTagIds,omitempty" validate:"dive gte=0"`
	PriceInHouseContactMin *int     `json:"priceInHouseContactMin,omitempty" validate:"gte=0"`
	PriceInHouseContactMax *int     `json:"priceInHouseContactMax,omitempty" validate:"gte=0"`
	PriceInHouseHourMin    *int     `json:"priceInHouseHourMin,omitempty" validate:"gte=0"`
	PriceInHouseHourMax    *int     `json:"priceInHouseHourMax,omitempty" validate:"gte=0"`
	PriceSaunaContactMin   *int     `json:"priceSaunaContactMin,omitempty" validate:"gte=0"`
	PriceSaunaContactMax   *int     `json:"priceSaunaContactMax,omitempty" validate:"gte=0"`
	PriceSaunaHourMin      *int     `json:"priceSaunaHourMin,omitempty" validate:"gte=0"`
	PriceSaunaHourMax      *int     `json:"priceSaunaHourMax,omitempty" validate:"gte=0"`
	PriceVisitContactMin   *int     `json:"priceVisitContactMin,omitempty" validate:"gte=0"`
	PriceVisitContactMax   *int     `json:"priceVisitContactMax,omitempty" validate:"gte=0"`
	PriceVisitHourMin      *int     `json:"priceVisitHourMin,omitempty" validate:"gte=0"`
	PriceVisitHourMax      *int     `json:"priceVisitHourMax,omitempty" validate:"gte=0"`
	PriceCarContactMin     *int     `json:"priceCarContactMin,omitempty" validate:"gte=0"`
	PriceCarContactMax     *int     `json:"priceCarContactMax,omitempty" validate:"gte=0"`
	PriceCarHourMin        *int     `json:"priceCarHourMin,omitempty" validate:"gte=0"`
	PriceCarHourMax        *int     `json:"priceCarHourMax,omitempty" validate:"gte=0"`
}

type ProfileResponse struct {
	ID                     string                   `json:"id"`
	UserID                 string                   `json:"userId"`
	Active                 bool                     `json:"active"`
	Phone                  string                   `json:"phone"`
	Name                   string                   `json:"name"`
	Age                    int                      `json:"age"`
	Height                 int                      `json:"height"`
	Weight                 int                      `json:"weight"`
	Bust                   float64                  `json:"bust"`
	Bio                    string                   `json:"bio"`
	AddressLatitude        string                   `json:"addressLatitude"`
	AddressLongitude       string                   `json:"addressLongitude"`
	CityID                 int                      `json:"cityId"`
	City                   *CityResponse            `json:"city"`
	BodyTypeID             *int                     `json:"bodyTypeId"`
	BodyType               *BodyTypeResponse        `json:"bodyTypeResponse"`
	EthnosID               *int                     `json:"ethnosId"`
	Ethnos                 *EthnosResponse          `json:"ethnos"`
	HairColorID            *int                     `json:"hairColorId"`
	HairColor              *HairColorResponse       `json:"hairColor"`
	IntimateHairCutID      *int                     `json:"intimateHairCutId"`
	IntimateHairCut        *IntimateHairCutResponse `json:"intimateHairCut"`
	PriceInHouseNightRatio float64                  `json:"priceInHouseNightRatio"`
	PriceInHouseContact    *int                     `json:"priceInHouseContact"`
	PriceInHouseHour       *int                     `json:"priceInHouseHour"`
	PriceSaunaNightRatio   float64                  `json:"priceSaunaNightRatio"`
	PriceSaunaContact      *int                     `json:"priceSaunaContact"`
	PriceSaunaHour         *int                     `json:"priceSaunaHour"`
	PriceVisitNightRatio   float64                  `json:"priceVisitNightRatio"`
	PriceVisitContact      *int                     `json:"priceVisitContact"`
	PriceVisitHour         *int                     `json:"priceVisitHour"`
	PriceCarNightRatio     float64                  `json:"priceCarNightRatio"`
	PriceCarContact        *int                     `json:"priceCarContact"`
	PriceCarHour           *int                     `json:"priceCarHour"`
	ContactPhone           string                   `json:"contactPhone"`
	ContactWA              string                   `json:"contactWA"`
	ContactTG              string                   `json:"contactTG"`
	Contacts               []ContactResponse        `json:"contacts"`
	Prices                 []PriceResponse          `json:"prices"`
	Moderated              bool                     `json:"moderated"`
	ModeratedAt            *time.Time               `json:"moderatedAt"`
	ModeratedBy            *uuid.UUID               `json:"moderatedBy"`
	Verified               bool                     `json:"verified"`
	VerifiedAt             *time.Time               `json:"verifiedAt"`
	VerifiedBy             *uuid.UUID               `json:"verifiedBy"`
	CreatedAt              time.Time                `json:"createdAt"`
	BodyArts               []ProfileBodyArtResponse `json:"bodyArts"`
	Photos                 []PhotoResponse          `json:"photos"`
	ProfileOptions         []ProfileOptionResponse  `json:"profileOptions"`
	Services               []ServiceResponse        `json:"services"`
	UpdatedBy              *uuid.UUID               `json:"updatedBy"`
}

type ContactResponse struct {
	ContactType string `json:"type"`
	Value       string `json:"value"`
}

type PriceResponse struct {
	TimeRange  string  `json:"timeRange"`
	Setting    string  `json:"setting"`
	Value      *int    `json:"value"`
	NightRatio float64 `json:"nightRatio"`
}
