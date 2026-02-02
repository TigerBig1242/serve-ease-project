package entities

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CategoryID   uint           `gorm:"primaryKey" json:"category_id"`
	CategoryName string         `json:"category_name"`
	Image        string         `json:"image"`
	IsAvailable  bool           `json:"is_available"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
