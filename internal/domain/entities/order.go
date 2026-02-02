package entities

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId       uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"order_id"`
	DiningTableID uint         `json:"dining_table_id"`
	DiningTable   DiningTable  `gorm:"references:TableID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"dining_table"`
	OrderItems    []OrderItems `gorm:"references:OrderId" json:"order_items_id"`
	OrderNumber   string       `json:"number"`
	TablePrice    float64      `json:"table_price"`
	Status        string       `json:"status"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type OrderItems struct {
	OrderItemsId uint      `gorm:"primaryKey" json:"order_items_id"`
	OrderId      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();" json:"order_id"`
	MenuId       uint      `json:"menu_id"`
	Quantity     int       `json:"quantity"`
	UnitPrice    int       `json:"unit_price"`
	Note         string    `json:"note"`
}

type OrderItemOptions struct {
	OrderItemOptionsId uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"order_item_options_is"`
	OrderItemsId       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();references:OrderItemsId" json:"order_items_id"`
	MenuOptionId       int       `gorm:"references:MenuOptionId"`
	PriceAtTime        float64   `json:"price_at_time"`
}
