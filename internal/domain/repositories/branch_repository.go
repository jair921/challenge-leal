package repositories

import "github.com/jair921/challenge-leal/internal/domain/entities"

// BranchRepository interfaz para operaciones relacionadas con sucursales
type BranchRepository interface {
	Create(branch *entities.Branch) error
	GetByID(id string) (*entities.Branch, error)
	List() ([]*entities.Branch, error)
}
