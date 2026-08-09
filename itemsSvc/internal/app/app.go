package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/config"
	grpcclient "github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/api/grpc/client"
	grpcserver "github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/api/grpc/server"
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/database"
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/repository"
	logger "github.com/AlekseiYuzhanin/GamePromoLogger"
)

func MustRun() error {
	log, err := logger.New(logger.NewConfig("objectservice", "info", ""))
	if err != nil {
		return err
	}
	log.Info("message", logger.Field{Key: "Test message", Value: "msg"})
	cfg := config.New("../../.env")
	db, err := database.New(cfg)
	if err != nil {
		return err
	}
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	client, err := grpcclient.New(cfg, log)
	if err != nil {
		return err
	}
	itemsRepository := repository.NewItemsRepository(log, db)
	srv, err := grpcserver.New(cfg, log, itemsRepository)
	if err != nil {
		return err
	}
	go func() {
		if err = srv.Start(); err != nil {
			log.Error("Error starting grpc server", logger.Field{Key: "err", Value: err})
		}
	}()
	<-sigint
	log.Info("Will gracefully shutdown")
	if err = db.Close(); err != nil {
		log.Error("Failed to close db", logger.Field{Key: "err", Value: err})
	}
	srv.Close()
	if err = client.Close(); err != nil {
		log.Error("Failed to close client", logger.Field{Key: "err", Value: err})
	}
	return nil
}
