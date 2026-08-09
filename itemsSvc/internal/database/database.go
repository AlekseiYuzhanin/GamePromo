package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/entities"
	_ "github.com/lib/pq"
	goose "github.com/pressly/goose/v3"
)

type Database interface {
	CreateSlot(context.Context, *entities.Slot) (*entities.Slot, error)
	Close() error
}
type database struct {
	db *sql.DB
}

func New(cfg *config.Config) (Database, error) {
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

	if err = goose.Up(db, "../../migrations"); err != nil {
		return nil, err
	}
	return &database{db: db}, nil
}

func (d *database) Close() error {
	err := goose.Down(d.db, "./migrations")
	if err != nil {
		return err
	}
	return d.db.Close()
}
