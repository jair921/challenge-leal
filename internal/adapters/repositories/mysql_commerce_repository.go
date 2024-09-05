package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

type MySQLCommerceRepository struct {
	db *sql.DB
}

func NewMySQLCommerceRepository(db *sql.DB) repositories.CommerceRepository {
	return &MySQLCommerceRepository{db: db}
}

func (r *MySQLCommerceRepository) GetCommerceByID(id string) (*entities.Commerce, error) {
	query := `SELECT id, name, branch_id FROM commerces WHERE id=?`
	commerce := &entities.Commerce{}
	err := r.db.QueryRow(query, id).Scan(&commerce.ID, &commerce.Name, &commerce.BranchID)
	if err != nil {
		return nil, err
	}
	return commerce, nil
}
