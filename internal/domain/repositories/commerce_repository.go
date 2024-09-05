package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// CommerceRepository define las operaciones de persistencia para los comercios.
type CommerceRepository interface {
	GetCommerceByID(id string) (*entities.Commerce, error)
}
