package entities

import (
	"time"
)

type DiningTable struct {
	TableID     uint   `gorm:"primaryKey" json:"table_id"`
	TableNumber string `json:"table_number"`
	QrToken     string `json:"qr_token"`
	Status      string `json:"status"`
	// OrderId     []Order   `gorm:"references:OrderId" json:"order_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
