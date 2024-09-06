package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
	"strconv"
)

type MySQLBranchRepository struct {
	db *sql.DB
}

// NewMySQLBranchRepository crea una nueva instancia del repositorio de sucursales
func NewMySQLBranchRepository(db *sql.DB) repositories.BranchRepository {
	return &MySQLBranchRepository{db: db}
}

// Create inserta una nueva sucursal en la base de datos
func (r *MySQLBranchRepository) Create(branch *entities.Branch) error {
	query := "INSERT INTO branches (commerce_id, name, address) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, branch.CommerceID, branch.Name, branch.Address)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	branch.ID = strconv.Itoa(int(id))

	return err
}

// GetByID obtiene una sucursal por su ID
func (r *MySQLBranchRepository) GetByID(id string) (*entities.Branch, error) {
	query := "SELECT id, commerce_id, name, address FROM branches WHERE id = ?"
	row := r.db.QueryRow(query, id)
	branch := &entities.Branch{}
	if err := row.Scan(&branch.ID, &branch.CommerceID, &branch.Name, &branch.Address); err != nil {
		return nil, err
	}
	return branch, nil
}

// List obtiene todas las sucursales
func (r *MySQLBranchRepository) List() ([]*entities.Branch, error) {
	query := "SELECT id, commerce_id, name, address FROM branches"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []*entities.Branch
	for rows.Next() {
		branch := &entities.Branch{}
		if err := rows.Scan(&branch.ID, &branch.CommerceID, &branch.Name, &branch.Address); err != nil {
			return nil, err
		}
		branches = append(branches, branch)
	}
	return branches, nil
}
