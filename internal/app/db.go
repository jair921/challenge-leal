package app

import (
	"database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
)

// InitDB inicializa la conexión a MySQL usando DSN
func InitDB(cfg *Config) (*sql.DB, error) {
	dsn := cfg.Database.DSN

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
