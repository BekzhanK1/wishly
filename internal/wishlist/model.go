package wishlist

import "time"

type Wishlist struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;notnull"`
	Title     string `gorm:"notnull"`
	Slug      string `gorm:"uniqueIndex;notnull"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
