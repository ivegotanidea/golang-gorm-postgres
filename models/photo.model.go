package models

import (
	"github.com/google/uuid"
	"time"
)

type Photo struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProfileID  uuid.UUID `gorm:"type:uuid;not null"`
	URL        string    `gorm:"type:varchar(255);not null"`
	PhrURL     string    `gorm:"type:varchar(255);null"`
	PreviewUrl string    `gorm:"type:varchar(255);null"`
	CreatedAt  time.Time `gorm:"type:timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp;null"`
	UpdatedBy  uuid.UUID `gorm:"type:uuid;null"`
	Hash       string    `gorm:"type:varchar(255);"`
	Disabled   bool      `gorm:"type:boolean;default:false"`
	Approved   bool      `gorm:"type:boolean;default:false"`
	Deleted    bool      `gorm:"type:boolean;default:false"`
}

type CreatePhotoRequest struct {
	URL string `json:"url" binding:"required" validate:"required,imageurl"`
}

type UpdatePhotoRequest struct {
	ID       string `json:"id" binding:"required" validate:"required,uuid"`
	Disabled bool   `json:"disabled" binding:"omitempty,oneof=true false"`
	Approved bool   `json:"approved" binding:"omitempty,oneof=true false"`
	Deleted  bool   `json:"deleted" binding:"omitempty,oneof=true false"`
}

type BulkUpdatePhotosRequest struct {
	Photos []UpdatePhotoRequest `json:"photos" binding:"required"`
}

type PhotoResponse struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	PhrURL     string `json:"phrUrl"`
	PreviewURL string `json:"previewUrl"`
	Disabled   bool   `json:"disabled"`
	Approved   bool   `json:"approved"`
	Deleted    bool   `json:"deleted"`
}
