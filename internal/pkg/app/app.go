package app

import (
	"log/slog"
	grpcapp "sso/internal/pkg/app/grpc"
	"sso/internal/services/auth"
	"sso/internal/storage/sqlite"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	gRPCServer := grpcapp.New(log, grpcPort, authService)
	return &App{
		GRPCServer: gRPCServer,
	}
}
