package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Amount      float64   `gorm:"type:decimal(10,2);not null"`
	Status      string    `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
	PaymentDate time.Time `gorm:"type:timestamp"`            // date at which payment is completed
	Type        string    `gorm:"type:varchar(50);not null"` // type: subscription, one_hour, three_hours, twelve_hours, two_days, one_week
}
