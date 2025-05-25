package wishlist

import "time"

type Wishlist struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index;notnull"`
	Title       string `gorm:"notnull"`
	Description string `gorm:"type:text"`
	ImageURL    string `gorm:"type:varchar(255)"`
	Slug        string `gorm:"uniqueIndex;notnull"`
	IsPublic    bool   `gorm:"notnull;default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Wishlist) TableName() string {
	return "wishlists"
}
