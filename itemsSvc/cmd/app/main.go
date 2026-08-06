package main

import (
	"github.com/AlekseiYuzhanin/GamePromo/itemsSvc/internal/app"
)

func main() {
	err := app.MustRun()
	if err != nil {
		panic(err)
	}
}
