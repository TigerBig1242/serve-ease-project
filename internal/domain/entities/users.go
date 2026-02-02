package entities

import (
	"time"
)

type Users struct {
	UserId         uint      `gorm:"primaryKey" json:"user_id"`
	Username       string    `json:"username"`
	PasswordHashed string    `json:"password_hashed"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
