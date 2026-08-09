package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/entities"
)

func (d *database) CreateSlot(ctx context.Context, slot *entities.Slot) (*entities.Slot, error) {
	var op string
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		op = "begin transaction"
		return nil, err
	}
	defer func() {
		if err != nil {
			fmt.Println("\n\n", err, op, slot.Type.Title, "\n\n")
			tx.Rollback()
		}
	}()
	slotTypeRows, err := tx.QueryContext(ctx, "SELECT id, title, created_at, updated_at FROM slots_type WHERE title = $1",
		slot.Type.Title)
	if err != nil {
		op = "query on slot type select"
		return nil, err
	}
	var slotType entities.SlotType
	if slotTypeRows.Next() {
		if err = slotTypeRows.Scan(&slotType.Id, &slotType.Title, &slotType.CreatedAt, &slotType.UpdatedAt); err != nil {
			op = "scan fields on slot type"
			return nil, err
		}
	} else {
		err = tx.QueryRowContext(ctx, "INSERT INTO slots_type(title) VALUES ($1) RETURNING id, title, created_at, updated_at",
			slot.Type.Title).
			Scan(&slotType.Id, &slotType.Title, &slotType.CreatedAt, &slotType.UpdatedAt)
		if err != nil {
			op = "insert into slot type"
			return nil, err
		}
	}
	slotTypeRows.Close()
	slotRarityRows, err := tx.QueryContext(ctx, "SELECT id, title, base_weight, multiplier, color_hex, created_at, updated_at "+
		"FROM slots_rarity WHERE title = $1",
		slot.Rarity.Title)
	if err != nil {
		op = "select on rarity rows"
		return nil, err
	}
	var slotRarity entities.SlotRarity
	if slotRarityRows.Next() {
		if err = slotRarityRows.Scan(&slotRarity.Id, &slotRarity.Title, &slotRarity.Weight, &slotRarity.Multiplier,
			&slotRarity.Color, &slotRarity.CreatedAt, &slotRarity.UpdatedAt); err != nil {
			op = "scan rows slot rarity"
			return nil, err
		}
	} else {
		err = tx.QueryRowContext(ctx,
			"INSERT INTO slots_rarity(title, base_weight, multiplier, color_hex) VALUES($1, $2, $3, $4) RETURNING id, title, "+
				"base_weight, multiplier, color_hex, created_at, updated_at",
			slot.Rarity.Title, slot.Rarity.Weight, slot.Rarity.Multiplier, slot.Rarity.Color).
			Scan(&slotRarity.Id, &slotRarity.Title, &slotRarity.Weight, &slotRarity.Multiplier,
				&slotRarity.Color, &slotRarity.CreatedAt, &slotRarity.UpdatedAt)
		if err != nil {
			op = "insert into slots rarity, new data"
			return nil, err
		}
	}
	slotRarityRows.Close()
	var res entities.Slot
	err = tx.QueryRowContext(ctx, "INSERT INTO slots(title, slot_type_id, slot_rarity_id) VALUES ($1,$2,$3) RETURNING id, title", slot.Title,
		slotType.Id, slotRarity.Id).Scan(&res.Id, &res.Title)
	if err != nil {
		op = "insert into slot table"
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		op = "commit failed"
		return nil, err
	}
	res.Type = slotType
	res.Rarity = slotRarity
	return &res, nil
}
