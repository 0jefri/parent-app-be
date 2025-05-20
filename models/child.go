package models

import "time"

type Child struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	ParentID  uint       `json:"parent_id"`
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	Parent Parent `gorm:"foreignKey:ParentID" json:"-"`
}
