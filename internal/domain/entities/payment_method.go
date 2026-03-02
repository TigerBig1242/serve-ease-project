package entities

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethods struct {
	PaymentTypeId uint   `gorm:"primaryKey" json:"payment_type_id"`
	PaymentMethod string `json:"payment_method"`
}

type Payment struct {
	PaymentId       uuid.UUID `gorm:"primaryKey" json:"payment_id"`
	OrderId         uuid.UUID `gorm:"references:OrderId" json:"order_id"`
	PaymentMethodId uint      `gorm:"references:PaymentTypeId" json:"payment_type_id"`
	Amount          float64   `json:"amount"`
	TransactionId   string    `json:"transaction_id"`
	Status          string    `json:"status"`
	PaidAt          time.Time `json:"paid-at"`
}
