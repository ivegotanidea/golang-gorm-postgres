package models

import (
	"time"

	"github.com/google/uuid"
)

type User2 struct {
	ID             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Phone          string    `gorm:"unique;size:50"`
	TelegramUserId int64     `gorm:"not null"`
	Password       string    `gorm:"not null"`
	Verified       bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
	Avatar         *string   `gorm:"type:varchar"`
}

type Profile struct {
	ID                  uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID              uuid.UUID `gorm:"foreignKey:UserID"`
	Active              bool      `gorm:"default:false"`
	Phone               string    `gorm:"unique;size:50"`
	Name                string    `gorm:"type:varchar(30)"`
	Age                 int       `gorm:"not null"`
	Height              int       `gorm:"not null"`
	Weight              int       `gorm:"not null"`
	Bust                float64   `gorm:"not null"`
	Ethnos              string    `gorm:"type:varchar(30)"`
	Bio                 string    `gorm:"type:varchar(1000)"`
	Moderated           bool      `gorm:"default:false"`
	Verified            bool      `gorm:"default:false"`
	PriceInHouseContact int
	PriceInHouseHour    int
	PriceSaunaContact   int
	PriceSaunaHour      int
	PriceVisitContact   int
	PriceVisitHour      int
	PriceCarContact     int
	PriceCarHour        int
	Latitude            float32
	Longitude           float32
	Photos              []Photo `gorm:"foreignKey:ProfileID"`
	ContactPhone        string  `gorm:"type:varchar(30)"`
	ContactWA           string  `gorm:"type:varchar(30)"`
	ContactTG           string  `gorm:"type:varchar(30)"`
}

// Photo model definition
type Photo struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProfileID uuid.UUID `gorm:"type:uuid"` // Foreign key to Profile
	URL       string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type Service struct {
	ID              uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID          uuid.UUID `gorm:"foreignKey:UserID"`
	ProfileID       uuid.UUID `gorm:"foreignKey:ProfileID"`
	CreatedAt       time.Time `gorm:"not null"`
	ProfileRatingID uuid.UUID `gorm:"type:uuid"`
	ProfileRating   Rating    `gorm:"foreignKey:ProfileRatingID"`

	UserRatingID uuid.UUID `gorm:"type:uuid"`
	UserRating   Rating    `gorm:"foreignKey:UserRatingID"`
}

type Rating struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ServiceID    uuid.UUID `gorm:"type:uuid;not null"` // Foreign key to Service
	LikedTags    []Tag     `gorm:"many2many:rating_liked_tags"`
	DislikedTags []Tag     `gorm:"many2many:rating_disliked_tags"`
	Review       string    `gorm:"type:text"`
	Score        int       `gorm:"not null;check:score >= 1 AND score <= 5"`
	CreatedAt    time.Time `gorm:"not null"`
}

// Tag model definition
type Tag struct {
	ID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar(30);unique;not null"`
}
