package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
	"strconv"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) repositories.UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) GetUserByID(id string) (*entities.User, error) {
	query := `SELECT id, name, email, points FROM users WHERE id=?`
	user := &entities.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Points)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLUserRepository) UpdateUser(user *entities.User) error {
	query := `UPDATE users SET points=? WHERE id=?`
	_, err := r.db.Exec(query, user.Points, user.ID)
	return err
}

// CreateUser inserta un nuevo usuario en la base de datos
func (r *MySQLUserRepository) CreateUser(user *entities.User) error {
	query := `INSERT INTO users (name, email, points) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, user.Name, user.Email, user.Points)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = strconv.Itoa(int(id))
	return err
}

// DeleteUser elimina un usuario de la base de datos por su ID
func (r *MySQLUserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *MySQLUserRepository) List() ([]*entities.User, error) {
	query := `SELECT id, name, email, points FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Points); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
