package wishlist

import (
	"time"

	"github.com/BekzhanK1/wishly/internal/user"
)

type CreateWishlistInput struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	IsPublic    bool   `json:"is_public"`
}

type WishlistOutput struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	Slug        string    `json:"slug"`
	IsPublic    bool      `json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WishlistOutputPublic struct {
	ID          uint              `json:"id"`
	UserID      uint              `json:"user_id"`
	Title       string            `json:"title"`
	Description string            `json:"description,omitempty"`
	ImageURL    string            `json:"image_url,omitempty"`
	Slug        string            `json:"slug"`
	IsPublic    bool              `json:"is_public"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	User        user.UserResponse `json:"user"`
}
