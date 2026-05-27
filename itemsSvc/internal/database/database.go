package database

import (
	"database/sql"
	"fmt"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func New(cfg *config.Config) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", cfg.Database.Host,
		cfg.Database.Port, cfg.Database.DbName, cfg.Database.User, cfg.Database.Password, cfg.Database.SslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}
