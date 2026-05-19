module github.com/AlekseiYuzhanin225/GamePromo/itemsSvc

go 1.26

require github.com/AlekseiYuzhanin225/GamePromo/logger v0.0.0

require (
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.28.0 // indirect
)

replace github.com/AlekseiYuzhanin225/GamePromo/logger => ../pkg/logger
