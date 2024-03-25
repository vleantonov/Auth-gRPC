package app

import (
	"log/slog"
	"sso/internal/config"
	"sso/internal/logger"
)

type App struct {
	cfg *config.Config
	log *slog.Logger
}

func New() *App {
	cfg, err := config.Load()
	if err != nil {
		panic(err.Error())
	}

	log, err := logger.New(cfg.Env)
	if err != nil {
		panic(err.Error())
	}

	return &App{
		cfg: cfg,
		log: log,
	}
}

func (a *App) Run() {
	a.log.Info("starting application")
}
