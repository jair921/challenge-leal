package services

import (
	"errors"
	"testing"

	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository es el mock del repositorio de usuarios
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) UpdateUser(user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserRepository) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserRepository) CreateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(id string) (*entities.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) List() ([]*entities.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

// Test para crear un usuario
func TestCreateUser(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	// Test case: successful user creation
	user := &entities.User{Name: "John Doe", Email: "john@example.com"}
	mockUserRepo.On("CreateUser", user).Return(nil)

	err := service.CreateUser(user)
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)

	// Test case: missing required fields
	invalidUser := &entities.User{}
	err = service.CreateUser(invalidUser)
	assert.EqualError(t, err, "name and email are required")
}

// Test para obtener un usuario por ID
func TestGetUserByID(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	// Test case: successful retrieval of a user
	user := &entities.User{ID: "1", Name: "John Doe", Email: "john@example.com"}
	mockUserRepo.On("GetUserByID", "1").Return(user, nil)

	result, err := service.GetUserByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockUserRepo.AssertExpectations(t)

	// Test case: error in retrieving a user
	mockUserRepo.On("GetUserByID", "2").Return(nil, errors.New("user not found"))

	result, err = service.GetUserByID("2")
	assert.Nil(t, result)
	assert.EqualError(t, err, "user not found")
	mockUserRepo.AssertExpectations(t)
}

// Test para listar todos los usuarios
func TestListUsers(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	service := NewUserService(mockUserRepo)

	// Test case: successful listing of users
	users := []*entities.User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
	}
	mockUserRepo.On("List").Return(users, nil)

	result, err := service.ListUsers()
	assert.NoError(t, err)
	assert.ElementsMatch(t, users, result)
	mockUserRepo.AssertExpectations(t)
}
