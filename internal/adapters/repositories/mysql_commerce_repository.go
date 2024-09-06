package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
	"strconv"
)

type MySQLCommerceRepository struct {
	db *sql.DB
}

// NewMySQLCommerceRepository crea una nueva instancia de MySQLCommerceRepository
func NewMySQLCommerceRepository(db *sql.DB) repositories.CommerceRepository {
	return &MySQLCommerceRepository{db: db}
}

// Create inserta un nuevo comercio en la base de datos
func (r *MySQLCommerceRepository) Create(commerce *entities.Commerce) error {
	query := `INSERT INTO commerces (name) VALUES (?)`
	result, err := r.db.Exec(query, commerce.Name)
	if err != nil {
		return err
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	commerce.ID = strconv.Itoa(int(id))

	return nil
}

// GetByID obtiene un comercio por su ID
func (r *MySQLCommerceRepository) GetByID(id string) (*entities.Commerce, error) {
	query := `SELECT id, name FROM commerces WHERE id=?`
	commerce := &entities.Commerce{}
	err := r.db.QueryRow(query, id).Scan(&commerce.ID, &commerce.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return commerce, nil
}

// List obtiene todos los comercios
func (r *MySQLCommerceRepository) List() ([]*entities.Commerce, error) {
	query := `SELECT id, name FROM commerces`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commerces []*entities.Commerce
	for rows.Next() {
		commerce := &entities.Commerce{}
		if err := rows.Scan(&commerce.ID, &commerce.Name); err != nil {
			return nil, err
		}
		commerces = append(commerces, commerce)
	}
	return commerces, nil
}
