package wishlist

import "errors"

type WishlistError struct {
	Code    string
	Message string
}

func (e WishlistError) Error() string {
	return e.Message
}

var (
	ErrInvalidWishlistInput         = errors.New("invalid wishlist input")
	ErrWishlistNotFound             = errors.New("wishlist not found")
	ErrWishlistCreationFailed       = errors.New("failed to create wishlist")
	ErrWishlistAlreadyExists        = errors.New("wishlist already exists")
	ErrWishlistSlugAlreadyUsed      = errors.New("wishlist slug is already used")
	ErrWishlistNotAuthorized        = errors.New("user not authorized to perform this action")
	ErrWishlistImageUploadFailed    = errors.New("failed to upload wishlist image")
	ErrWishlistDeletionFailed       = errors.New("failed to delete wishlist")
	ErrWishlistUpdateFailed         = errors.New("failed to update wishlist")
	ErrWishlistSlugGenerationFailed = errors.New("failed to generate unique wishlist slug")
	ErrWishlistPublicAccessDenied   = errors.New("public access to wishlist denied")
)
