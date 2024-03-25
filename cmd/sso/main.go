package main

import (
	"sso/internal/config"
	"sso/internal/logger"
	"sso/internal/pkg/app"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err.Error())
	}

	log, err := logger.New(cfg.Env)
	if err != nil {
		panic(err.Error())
	}

	a := app.New(
		log,
		cfg.GRPC.Port,
		cfg.StoragePath,
		cfg.TokenTTL,
	)
	a.GRPCServer.MustRun()
}
