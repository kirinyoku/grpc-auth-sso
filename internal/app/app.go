package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/kirinyoku/grpc-auth-sso/internal/app/grpc"
	"github.com/kirinyoku/grpc-auth-sso/internal/services/auth"
	"github.com/kirinyoku/grpc-auth-sso/internal/storage/sqlite"
)

type App struct {
	GRPCApp *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCApp: grpcApp,
	}
}
