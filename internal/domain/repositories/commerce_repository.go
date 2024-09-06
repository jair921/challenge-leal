package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// CommerceRepository define las operaciones de persistencia para los comercios.
type CommerceRepository interface {
	Create(commerce *entities.Commerce) error
	GetByID(id string) (*entities.Commerce, error)
	List() ([]*entities.Commerce, error)
}
