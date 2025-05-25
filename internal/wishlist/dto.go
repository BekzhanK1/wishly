package wishlist

import (
	"time"

	"github.com/BekzhanK1/wishly/internal/user"
)

type CreateWishlistInput struct {
	Title string `json:"title"`
}

type WishlistOutput struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WishlistOutputPublic struct {
	ID                uint      `json:"id"`
	UserID            uint      `json:"user_id"`
	Title             string    `json:"title"`
	Slug              string    `json:"slug"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	user.UserResponse `json:"user"`
}
