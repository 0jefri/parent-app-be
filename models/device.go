package models

import "time"

type Device struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ChildID    uint   `json:"child_id"`
	DeviceName string `json:"device_name"`
	Status     string `json:"status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Child Child `gorm:"foreignKey:ChildID" json:"-"`
}
