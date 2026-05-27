package main

import (
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/database"
	logger "github.com/AlekseiYuzhanin/GamePromoLogger"
)

func main() {
	log, err := logger.New(logger.NewConfig("objectservice", "info", ""))
	if err != nil {
		panic(err)
	}
	log.Info("message", logger.Field{Key: "Test message", Value: "msg"})
	cfg := config.New("../.env")
	_, err = database.New(cfg)
	if err != nil {
		panic(err)
	}

}
