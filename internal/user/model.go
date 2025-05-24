package user

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string // Password hash
	CreatedAt time.Time
	UpdatedAt time.Time
}
