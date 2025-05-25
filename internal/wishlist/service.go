package wishlist

import (
	"github.com/BekzhanK1/wishly/internal/user"
	"github.com/BekzhanK1/wishly/pkg/slug"
)

type Service interface {
	CreateWishlist(*CreateWishlistInput, uint) error
	GetWishlistByID(uint) (*WishlistOutput, error)
	GetWishlistBySlug(string) (*WishlistOutput, error)
	GetAllWishlistsByUser(uint) ([]*WishlistOutput, error)
	GetWishlistByUsernameAndSlug(string, string) (*WishlistOutputPublic, error)
	UpdateWishlist(*Wishlist) error
	DeleteWishlist(uint, uint) error
	IsSlugUsedByUser(string, uint) bool
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateWishlist(input *CreateWishlistInput, userID uint) error {
	if input == nil {
		return ErrInvalidWishlistInput
	}

	slug := slug.GenerateUniqueSlug(input.Title, func(slug string) bool {
		return s.repo.SlugExistsForUser(slug, userID)
	})

	wishlist := &Wishlist{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		IsPublic:    input.IsPublic,
		Slug:        slug,
		UserID:      userID,
	}

	if err := s.repo.Create(wishlist); err != nil {
		return ErrWishlistCreationFailed
	}
	return nil
}

func (s *service) GetWishlistByID(id uint) (*WishlistOutput, error) {
	wishlist, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrWishlistNotFound
	}

	wishlistOutput := toOutput(wishlist)

	return wishlistOutput, nil
}

func (s *service) GetWishlistByUsernameAndSlug(username string, slug string) (*WishlistOutputPublic, error) {
	wishlist, user, err := s.repo.FindByUsernameAndSlug(username, slug)
	if err != nil {
		return nil, ErrWishlistNotFound
	}

	if !wishlist.IsPublic {
		return nil, ErrWishlistPublicAccessDenied
	}

	wishlistOutput := toOutputWithUser(wishlist, user)
	return wishlistOutput, nil
}

func (s *service) GetWishlistBySlug(slug string) (*WishlistOutput, error) {
	wishlist, err := s.repo.FindBySlug(slug)
	if err != nil {
		return nil, ErrWishlistNotFound
	}

	wishlistOutput := toOutput(wishlist)

	return wishlistOutput, nil
}

func (s *service) GetAllWishlistsByUser(userID uint) ([]*WishlistOutput, error) {
	wishlists, err := s.repo.FindAllByUser(userID)
	if err != nil {
		return nil, ErrWishlistNotFound
	}
	wishlistOutputs := make([]*WishlistOutput, len(wishlists))
	for i, wishlist := range wishlists {
		wishlistOutputs[i] = toOutput(wishlist)
	}
	return wishlistOutputs, nil
}

func (s *service) UpdateWishlist(w *Wishlist) error {
	if w == nil {
		return ErrInvalidWishlistInput
	}

	if s.repo.SlugExistsForUser(w.Slug, w.UserID) {
		w.Slug = slug.GenerateUniqueSlug(w.Slug, func(slug string) bool {
			return s.repo.SlugExistsForUser(slug, w.UserID)
		})
	}

	if err := s.repo.Update(w); err != nil {
		return ErrWishlistUpdateFailed
	}
	return nil
}

func (s *service) DeleteWishlist(id uint, userID uint) error {
	if err := s.repo.Delete(id, userID); err != nil {
		return ErrWishlistDeletionFailed
	}
	return nil
}

func (s *service) IsSlugUsedByUser(slug string, userID uint) bool {
	return s.repo.SlugExistsForUser(slug, userID)
}

func toOutput(w *Wishlist) *WishlistOutput {
	return &WishlistOutput{
		ID:          w.ID,
		Title:       w.Title,
		Description: w.Description,
		ImageURL:    w.ImageURL,
		Slug:        w.Slug,
		IsPublic:    w.IsPublic,
		UserID:      w.UserID,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}
}

func toOutputWithUser(w *Wishlist, u *user.User) *WishlistOutputPublic {
	return &WishlistOutputPublic{
		ID:          w.ID,
		Title:       w.Title,
		Description: w.Description,
		ImageURL:    w.ImageURL,
		Slug:        w.Slug,
		IsPublic:    w.IsPublic,
		UserID:      w.UserID,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
		User:        *user.ToUserResponse(u),
	}
}
