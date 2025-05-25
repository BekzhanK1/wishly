package wishlist

import (
	"github.com/BekzhanK1/wishly/internal/user"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*Wishlist) error
	FindByID(uint) (*Wishlist, error)
	FindBySlug(string) (*Wishlist, error)
	FindByUsernameAndSlug(string, string) (*Wishlist, *user.User, error)
	FindAllByUser(uint) ([]*Wishlist, error)
	Update(*Wishlist) error
	Delete(uint, uint) error
	SlugExistsForUser(string, uint) bool
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(wishlist *Wishlist) error {
	return r.db.Create(wishlist).Error
}

func (r *repository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Wishlist{}).Error
}

func (r *repository) FindByUsernameAndSlug(username string, slug string) (*Wishlist, *user.User, error) {
	var wishlist Wishlist
	err := r.db.Where("slug = ? AND user_id = (SELECT id FROM users WHERE username = ?)", slug, username).First(&wishlist).Error
	if err != nil {
		return nil, nil, err
	}
	var user user.User
	err = r.db.Where("id = ?", wishlist.UserID).First(&user).Error
	if err != nil {
		return nil, nil, err
	}
	return &wishlist, &user, nil
}

func (r *repository) FindAllByUser(userID uint) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	err := r.db.Where("user_id = ?", userID).Find(&wishlists).Error
	return wishlists, err
}

func (r *repository) FindByID(id uint) (*Wishlist, error) {
	var wishlist Wishlist
	err := r.db.Where("id = ?", id).First(&wishlist).Error
	return &wishlist, err
}

func (r *repository) FindBySlug(slug string) (*Wishlist, error) {
	var wishlist Wishlist
	err := r.db.Where("slug = ?", slug).First(&wishlist).Error
	return &wishlist, err
}

func (r *repository) SlugExistsForUser(slug string, userID uint) bool {
	var count int64
	r.db.Model(&Wishlist{}).Where("slug = ? AND user_id = ?", slug, userID).Count(&count)
	return count > 0
}

func (r *repository) Update(wishlist *Wishlist) error {
	return r.db.Save(wishlist).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
