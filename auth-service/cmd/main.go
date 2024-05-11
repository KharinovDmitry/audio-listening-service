package main

import (
	"auth-service/cmd/migrator"
	"auth-service/config"
	_ "auth-service/docs"
	"auth-service/internal/server/rest/router"
	"auth-service/internal/service"
	storage2 "auth-service/internal/storage"
	"auth-service/lib/adapter/messageBroker/rabbitMQ"
	"flag"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @description auth-service API
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

	migrator.Run(cfg.DBConnStr, cfg.MigrationsPath)

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
		panic("connect to message broker error: " + err.Error())
	}
	defer func() {
		if err = broker.Close(); err != nil {
			panic("close message broker error: " + err.Error())
		}
	}()

	services := service.NewManager(storage, broker, cfg.Env, cfg.JwtSecret, cfg.Salt, cfg.TokenTTL)

	services.Logger.Info("auth-service started")
	err = router.Run(cfg.Http.Port, services, storage)
	if err != nil {
		panic("start server error: " + err.Error())
	}
}
