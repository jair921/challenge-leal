package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// TransactionRepository define las operaciones de persistencia para las transacciones.
type TransactionRepository interface {
	AccumulatePoints(transaction *entities.Transaction) error
	GetTransactionsByUser(userID string) ([]*entities.Transaction, error)
}
