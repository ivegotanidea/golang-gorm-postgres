package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string    `gorm:"type:varchar(20);not null"`
	Phone          string    `gorm:"type:varchar(30);uniqueIndex"`
	TelegramUserId int64     `gorm:"type:bigint;not null;uniqueIndex"`
	Password       string    `gorm:"type:varchar(255);not null"`
	Active         bool      `gorm:"type:boolean;default:true"`
	Verified       bool      `gorm:"type:boolean;default:false"`
	CreatedAt      time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt      time.Time `gorm:"type:timestamp;not null"`
	Avatar         string    `gorm:"type:varchar(255)"`
	HasProfile     bool      `gorm:"type:boolean"`
	Profiles       []Profile `gorm:"foreignKey:UserID"`
	Services       []Service `gorm:"foreignKey:ClientUserID"`
	Tier           string    `gorm:"type:varchar(50);not null;default:'basic'"`
}

type SignUpInput struct {
	Phone           string `json:"phone" binding:"required,min=11,max=11"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,min=8"`
}

type SignInInput struct {
	Phone    string `json:"phone" binding:"required,min=11,max=11"`
	Password string `json:"password" binding:"required,min=8"`
}

type BotSignUpInput struct {
	Name           string `json:"name" binding:"required,min=5"`
	Phone          string `json:"phone" binding:"required,min=11,max=11"`
	TelegramUserId string `json:"telegramUserId" binding:"required"`
}

type BotSignInInput struct {
	TelegramUserId string `json:"telegramUserId"  binding:"required"`
}

type FindUserQuery struct {
	Id             string `form:"id"`
	Phone          string `form:"phone"`
	TelegramUserId int64  `form:"telegramUserId"`
}

type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	TelegramUserID int64     `json:"telegramUserId"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Password       string    `json:"password,omitempty"`
	Avatar         string    `json:"photo,omitempty"`
	Verified       bool      `json:"verified"`
	Active         bool      `json:"active"`
	Tier           string    `json:"tier"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UpdateUserPrivileged struct {
	Name           string `json:"name,omitempty" validate:"omitempty,min=3,max=20"`
	Phone          string `json:"phone,omitempty" validate:"omitempty,min=11,max=11"`
	TelegramUserId string `json:"telegramUserId,omitempty" validate:"omitempty,min=6"`
	Avatar         string `json:"photo,omitempty"  validate:"omitempty,imageurl"`
	Verified       bool   `json:"verified,omitempty" validate:"omitempty,boolean"`
	Tier           string `json:"tier,omitempty" validate:"omitempty,oneof=basic experienced guru moderator admin"`
	Active         bool   `json:"active,omitempty" validate:"omitempty,boolean"`
}

type UpdateUser struct {
	Name   string `json:"name,omitempty" validate:"omitempty,min=3,max=20"`
	Phone  string `json:"phone,omitempty" validate:"omitempty,min=11,max=11"`
	Avatar string `json:"photo,omitempty"  validate:"omitempty,imageurl"`
}

type UpdateUserPassword struct {
	OldPassword string `json:"password" binding:"required,min=8"`
	NewPassword string `json:"newPassword" binding:"required,min=8"`
}
