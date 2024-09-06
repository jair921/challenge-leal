package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// TransactionRepository define las operaciones de persistencia para las transacciones.
type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	GetByUserID(userID string) ([]*entities.Transaction, error)
}
