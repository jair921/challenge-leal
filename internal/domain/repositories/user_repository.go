package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// UserRepository define las operaciones de persistencia para los usuarios.
type UserRepository interface {
	GetUserByID(id string) (*entities.User, error)
	UpdateUser(user *entities.User) error
}
