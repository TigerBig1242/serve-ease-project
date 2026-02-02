package entities

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	MenuId      uint           `gorm:"primaryKey" json:"menu_id"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `gorm:"references:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"category"`
	MenuName    string         `json:"menu_name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Image       string         `json:"image"`
	IsAvailable bool           `json:"is_available"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type MenuOptions struct {
	MenuOptionId int     `gorm:"primaryKey" json:"menu_option_id"`
	MenuId       uint    `gorm:"references:MenuId" json:"menu_id"`
	GroupName    string  `json:"group_name"`
	Label        string  `json:"label"`
	ExtraPrice   float64 `json:"extra_price"`
}
