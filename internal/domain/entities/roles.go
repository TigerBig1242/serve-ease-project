package entities

import (
	"time"
)

type Roles struct {
	RoleId    uint      `gorm:"primaryKey" json:"role_id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}
