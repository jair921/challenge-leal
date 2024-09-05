package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
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
