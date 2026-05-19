module github.com/AlekseiYuzhanin/GamePromo/itemsSvc

replace github.com/AlekseiYuzhanin/GamePromo/logger => ../pkg/logger

go 1.26

require (
	github.com/AlekseiYuzhanin/GamePromo/logger v0.0.0-00010101000000-000000000000 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.28.0 // indirect
)
