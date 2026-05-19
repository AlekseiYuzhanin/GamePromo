package main

import "github.com/AlekseiYuzhanin/GamePromo/logger"

func main() {
	loggerCfg := logger.NewConfig("itemsSvc", "debug", "")
	log, err := logger.New(loggerCfg)
	if err != nil {
		return
	}
	log.Info("Some kind of message")
}
