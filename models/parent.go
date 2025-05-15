package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Parent struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100"`
	Email       string `gorm:"size:100;uniqueIndex"`
	PhoneNumber string `gorm:"size:20"`
	FirebaseUID string `gorm:"size:128;uniqueIndex;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
