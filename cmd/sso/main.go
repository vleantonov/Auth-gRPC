package main

import (
	"os"
	"os/signal"
	"sso/internal/config"
	"sso/internal/logger"
	"sso/internal/pkg/app"
	"syscall"
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
	go a.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	a.GRPCServer.Stop()
}
