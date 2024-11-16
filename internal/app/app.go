package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/kirinyoku/grpc-auth-sso/internal/app/grpc"
)

type App struct {
	GRPCApp *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokneTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCApp: grpcApp,
	}
}
