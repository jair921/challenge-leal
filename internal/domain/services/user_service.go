package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

// UserService define los métodos relacionados con la lógica de negocio de usuarios.
type UserService interface {
	CreateUser(user *entities.User) error
	GetUserByID(id string) (*entities.User, error)
	ListUsers() ([]*entities.User, error)
}

// userService es una implementación concreta de UserService
type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService crea un nuevo servicio de usuarios
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser crea un nuevo usuario
func (s *userService) CreateUser(user *entities.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return s.userRepo.CreateUser(user)
}

// GetUserByID obtiene un usuario por su ID
func (s *userService) GetUserByID(id string) (*entities.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// ListUsers obtiene todos los usuarios
func (s *userService) ListUsers() ([]*entities.User, error) {
	return s.userRepo.List()
}
