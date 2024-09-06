package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// UserRepository define las operaciones de persistencia para los usuarios.
type UserRepository interface {
	GetUserByID(id string) (*entities.User, error) // Obtener un usuario por su ID
	UpdateUser(user *entities.User) error          // Actualizar un usuario
	CreateUser(user *entities.User) error          // Crear un nuevo usuario
	DeleteUser(id string) error                    // Eliminar un usuario
	List() ([]*entities.User, error)               // Listar usuarios
}
