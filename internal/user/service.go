package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Register(input RegisterInput) (*User, error)
	ValidateCredentials(input LoginInput) (*User, error)
}

type service struct {
	repo Repository
}

func (s *service) ValidateCredentials(input LoginInput) (*User, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if hashErr != nil {
		return nil, hashErr

	}

	return user, nil
}

func (s *service) Register(input RegisterInput) (*User, error) {
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
	return user, err

}

func NewService(repo Repository) Service {
	return &service{repo}
}
