package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Phone          string    `gorm:"type:varchar(255)"`
	TelegramUserId int64     `gorm:"type:bigint;not null"`
	Password       string    `gorm:"type:varchar(42);not null"`
	Verified       bool      `gorm:"type:boolean"`
	CreatedAt      time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt      time.Time `gorm:"type:timestamp;not null"`
	Avatar         string    `gorm:"type:varchar(255)"`
	HasProfile     bool      `gorm:"type:boolean"`
	Profiles       []Profile `gorm:"foreignKey:UserID"`
	Services       []Service `gorm:"foreignKey:ClientUserID"`
}
