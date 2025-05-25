package item

import (
	"time"
)

type Reservation struct {
	ID         uint      `gorm:"primaryKey"`
	ItemID     uint      `gorm:"not null;uniqueIndex"`
	UserID     uint      `gorm:"not null"`
	ReservedAt time.Time `gorm:"autoCreateTime"`

	Item Item `gorm:"constraint:OnDelete:CASCADE;"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	WishlistID  uint   `gorm:"not null;index"`
	Title       string `gorm:"not null"`
	Description string
	URL         string
	ImageURL    string
	Price       float64 `gorm:"not null;default:0"`
	IsReserved  bool    `gorm:"not null;default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Reservation *Reservation `gorm:"foreignKey:ItemID;constraint:OnDelete:SET NULL;"`
}
