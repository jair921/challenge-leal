package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

type MySQLTransactionRepository struct {
	db *sql.DB
}

func NewMySQLTransactionRepository(db *sql.DB) repositories.TransactionRepository {
	return &MySQLTransactionRepository{db: db}
}

func (r *MySQLTransactionRepository) AccumulatePoints(transaction *entities.Transaction) error {
	query := `INSERT INTO transactions (id, user_id, commerce_id, branch_id, amount, points_earned, cashback_earned) 
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, transaction.ID, transaction.UserID, transaction.CommerceID, transaction.BranchID, transaction.Amount, transaction.PointsEarned, transaction.CashbackEarned)
	return err
}

func (r *MySQLTransactionRepository) GetTransactionsByUser(userID string) ([]*entities.Transaction, error) {
	query := `SELECT id, user_id, commerce_id, branch_id, amount, points_earned, cashback_earned 
	          FROM transactions WHERE user_id=?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*entities.Transaction
	for rows.Next() {
		transaction := &entities.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.CommerceID, &transaction.BranchID, &transaction.Amount, &transaction.PointsEarned, &transaction.CashbackEarned)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
