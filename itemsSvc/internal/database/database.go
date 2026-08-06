package database

import (
	"database/sql"
	"fmt"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	_ "github.com/lib/pq"
	goose "github.com/pressly/goose/v3"
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

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if err = goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err = goose.Up(db, "./migrations"); err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	err := goose.Down(d.db, "./migrations")
	if err != nil {
		return err
	}
	return d.db.Close()
}
