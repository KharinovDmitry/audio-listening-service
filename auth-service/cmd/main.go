package main

import (
	"auth-service/config"
	logger2 "auth-service/internal/service"
	"flag"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "path", "", "path to config file")
	flag.Parse()

	cfg, err := config.Load(cfgPath)
	if err != nil {
		panic("Ошибка при чтении конфига: " + err.Error())
	}

	logger := logger2.NewLogger(cfg.Env)
}
