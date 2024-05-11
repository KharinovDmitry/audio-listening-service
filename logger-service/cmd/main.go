package main

import (
	"context"
	"flag"
	"logger-service/cmd/migrator"
	"logger-service/config"
	"logger-service/internal/service"
	storage2 "logger-service/internal/storage"
	"logger-service/lib/adapter/messageBroker/rabbitMQ"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "path", "", "path to config file")
	flag.Parse()
	if cfgPath == "" {
		panic("config file path is required")
	}

	cfg, err := config.Load(cfgPath)
	if err != nil {
		panic("read config error: " + err.Error())
	}

	migrator.MustRun(cfg.DBConnStr, cfg.MigrationsPath)

	storage := storage2.NewStorage()
	if err = storage.Init(cfg.TimeoutDB, cfg.DBConnStr); err != nil {
		panic("open db connection error: " + err.Error())
	}
	defer func() {
		if err = storage.Close(); err != nil {
			panic("close db connection error: " + err.Error())
		}
	}()

	broker := rabbitMQ.NewRabbitMQAdapter()
	if err = broker.Connect(cfg.MBConnStr); err != nil {
		panic("connect broker error: " + err.Error())
	}
	defer func() {
		if err = broker.Close(); err != nil {
			panic("close broker connection error: " + err.Error())
		}
	}()

	services := service.NewManager(storage, broker)

	if err = services.Logger.StartWriteLogs(context.Background()); err != nil {
		panic("start write logs error: " + err.Error())
	}
}
