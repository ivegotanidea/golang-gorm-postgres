package models

import (
	"github.com/google/uuid"
	"time"
)

type Photo struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProfileID  uuid.UUID `gorm:"type:uuid;not null"`
	Extension  string    `gorm:"type:varchar(10);"`
	URL        string    `gorm:"type:varchar(255);not null"`
	PhrURL     string    `gorm:"type:varchar(255);null"`
	PreviewUrl string    `gorm:"type:varchar(255);null"`
	CreatedAt  time.Time `gorm:"type:timestamp"`
	Disabled   bool      `gorm:"type:boolean;default:false"`
	Approved   bool      `gorm:"type:boolean;default:false"`
	Deleted    bool      `gorm:"type:boolean;default:false"`
}

type CreatePhotoRequest struct {
	URL string `json:"url" binding:"required" validate:"required,imageurl"`
}

type PhotoResponse struct {
	URL        string `json:"url"`
	PhrURL     string `json:"phrUrl"`
	PreviewURL string `json:"previewUrl"`
	Disabled   bool   `json:"disabled"`
	Approved   bool   `json:"approved"`
	Deleted    bool   `json:"deleted"`
}
