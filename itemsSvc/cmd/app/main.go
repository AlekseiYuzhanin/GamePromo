package main

import "github.com/AlekseiYuzhanin225/GamePromo/logger"

func main() {
	loggerCfg := logger.NewConfig("objectservice", "info", "")
	log, err := logger.New(loggerCfg)
	if err != nil {
		panic(err)
	}
	log.Info("message", logger.Field{Key: "Test message", Value: "msg"})
}
