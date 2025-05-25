package user

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidInput        = errors.New("invalid input")
	ErrEmailAlreadyUsed    = errors.New("email is already used")
	ErrUsernameAlreadyUsed = errors.New("username is already used")
	ErrPasswordTooShort    = errors.New("password must be at least 8 characters long")
	ErrInvalidEmailFormat  = errors.New("invalid email format")
	ErrUserNotAuthorized   = errors.New("user not authorized")
	ErrCouldntCreateUser   = errors.New("could not create user")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)
