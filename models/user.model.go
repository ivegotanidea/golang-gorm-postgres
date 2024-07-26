package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string    `gorm:"type:varchar(20);not null"`
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

type SignUpInput struct {
	Phone           string `json:"phone" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,min=8"`
}

type SignInInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type BotSignUpInput struct {
	Name           string `json:"name" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	TelegramUserId string `json:"email" binding:"required"`
}

type BotSignInInput struct {
	TelegramUserId string `json:"telegramUserId"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password,omitempty"`
	Avatar    string    `json:"photo,omitempty"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
