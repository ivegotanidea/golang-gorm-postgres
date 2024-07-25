package models

import (
	"github.com/google/uuid"
)

type UserTag struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"type:varchar(255)"`
}
