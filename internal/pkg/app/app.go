package app

import (
	"log/slog"
	grpcapp "sso/internal/pkg/app/grpc"
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
	gRPCServer := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServer: gRPCServer,
	}
}
