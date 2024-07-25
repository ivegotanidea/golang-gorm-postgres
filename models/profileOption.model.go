package models

import (
	"github.com/google/uuid"
)

type ProfileOption struct {
	ProfileID    uuid.UUID  `gorm:"primaryKey;type:uuid"`
	ProfileTagID uuid.UUID  `gorm:"primaryKey;type:uuid"`
	Price        int64      `gorm:"type:int64"`
	ProfileTag   ProfileTag `gorm:"foreignKey:ProfileTagID"`
}
