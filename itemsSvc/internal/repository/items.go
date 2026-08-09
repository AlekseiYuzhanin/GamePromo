package repository

import (
	"context"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/database"
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/entities"
	logger "github.com/AlekseiYuzhanin/GamePromoLogger"
)

type ItemsRepository interface {
	CreateSlot(ctx context.Context, slot *entities.Slot) (*entities.Slot, error)
}

type itemsRepository struct {
	db  database.Database
	log logger.Logger
}

func NewItemsRepository(log logger.Logger, db database.Database) ItemsRepository {
	return &itemsRepository{
		db:  db,
		log: log,
	}
}

func (i *itemsRepository) CreateSlot(ctx context.Context, slot *entities.Slot) (*entities.Slot, error) {
	res, err := i.db.CreateSlot(ctx, slot)
	if err != nil {
		i.log.Error("Failed to create slot", logger.Field{Key: "err", Value: err})
		return nil, err
	}
	return res, nil
}
