package database

import (
	"database/sql"
	"fmt"
	"github.com/gui-laranjeira/livreria/configs"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection(cfg *configs.DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
		return nil, err
	}
	return db, nil
}
