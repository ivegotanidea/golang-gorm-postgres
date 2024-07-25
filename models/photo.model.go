package models

import (
	"github.com/google/uuid"
	"time"
)

type Photo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProfileID uuid.UUID `gorm:"type:uuid;not null"`
	URL       string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	Disabled  bool      `gorm:"type:boolean"`
	Deleted   bool      `gorm:"type:boolean"`
}
