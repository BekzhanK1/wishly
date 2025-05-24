package user_test

import (
	"errors"
	"testing"

	"github.com/BekzhanK1/wishly/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepo mocks the user.Repository interface
type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) FindByEmail(email string) (*user.User, error) {
	args := m.Called(email)
	if u := args.Get(0); u != nil {
		return u.(*user.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockRepo) FindByID(id uint) (*user.User, error) {
	args := m.Called(id)
	if u := args.Get(0); u != nil {
		return u.(*user.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockRepo) Create(u *user.User) error {
	args := m.Called(u)
	return args.Error(0)
}

// --- Tests ---

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	input := user.RegisterInput{
		Username: "bekzhan",
		Email:    "bek@example.com",
		Password: "securepass",
	}

	mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(nil)

	resp, err := svc.Register(input)

	assert.NoError(t, err)
	assert.Equal(t, input.Username, resp.Username)
	assert.Equal(t, input.Email, resp.Email)
	mockRepo.AssertExpectations(t)
}

func TestRegister_RepoFails(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	input := user.RegisterInput{
		Username: "fail",
		Email:    "fail@example.com",
		Password: "1234",
	}

	mockRepo.On("Create", mock.Anything).Return(errors.New("db error"))

	resp, err := svc.Register(input)

	assert.Nil(t, resp)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMe_Success(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	expectedUser := &user.User{
		ID:       1,
		Username: "bekzhan",
		Email:    "bek@example.com",
	}

	mockRepo.On("FindByID", uint(1)).Return(expectedUser, nil)

	resp, err := svc.Me(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Username, resp.Username)
	assert.Equal(t, expectedUser.Email, resp.Email)
	mockRepo.AssertExpectations(t)
}

func TestMe_UserNotFound(t *testing.T) {
	mockRepo := new(mockRepo)
	svc := user.NewService(mockRepo)

	mockRepo.On("FindByID", uint(404)).Return(nil, errors.New("not found"))

	resp, err := svc.Me(404)

	assert.Nil(t, resp)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
