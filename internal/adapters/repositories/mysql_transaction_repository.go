package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
	"strconv"
)

// MySQLTransactionRepository es la implementación concreta de TransactionRepository para MySQL
type MySQLTransactionRepository struct {
	db *sql.DB
}

// NewMySQLTransactionRepository crea una nueva instancia de MySQLTransactionRepository
func NewMySQLTransactionRepository(db *sql.DB) repositories.TransactionRepository {
	return &MySQLTransactionRepository{db: db}
}

// Create inserta una nueva transacción en la base de datos
func (r *MySQLTransactionRepository) Create(transaction *entities.Transaction) error {
	query := `
		INSERT INTO transactions (user_id, commerce_id, branch_id, campaign_id, amount, points_earned, cashback_earned, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, transaction.UserID, transaction.CommerceID, transaction.BranchID, transaction.CampaignID, transaction.Amount, transaction.PointsEarned, transaction.CashbackEarned, transaction.CreatedAt)

	if err != nil {
		return err
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	transaction.ID = strconv.Itoa(int(id))
	return err
}

// GetByUserID obtiene todas las transacciones de un usuario específico
func (r *MySQLTransactionRepository) GetByUserID(userID string) ([]*entities.Transaction, error) {
	query := `SELECT id, user_id, commerce_id, branch_id, campaign_id, amount, points_earned, cashback_earned, created_at FROM transactions WHERE user_id=?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*entities.Transaction
	for rows.Next() {
		transaction := &entities.Transaction{}
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.CommerceID, &transaction.BranchID, &transaction.CampaignID, &transaction.Amount, &transaction.PointsEarned, &transaction.CashbackEarned, &transaction.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}
