package user

import (
	"github.com/BekzhanK1/wishly/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ValidateCredentials(LoginInput) (*LoginOutput, error)
	Register(RegisterInput) (*UserResponse, error)
	Me(uint) (*UserResponse, error)
}

type service struct {
	repo Repository
}

func (s *service) ValidateCredentials(input LoginInput) (*LoginOutput, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if hashErr != nil {
		return nil, hashErr

	}

	accessToken, refreshToken, err := auth.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return toLoginResponse(user, accessToken, refreshToken), nil
}

func (s *service) Register(input RegisterInput) (*UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user), err

}

func (s *service) Me(userID uint) (*UserResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func toUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func toLoginResponse(u *User, accessToken string, refreshToken string) *LoginOutput {
	return &LoginOutput{
		UserResponse: UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
