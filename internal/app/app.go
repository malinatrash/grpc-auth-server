package app

import (
	grpcapp "github.com/malinatrash/grpc-auth-server/internal/app/grpc"
	"github.com/malinatrash/grpc-auth-server/internal/services/auth"
	"github.com/malinatrash/grpc-auth-server/internal/services/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}
	authService := auth.New(log, storage, storage, storage, tokenTTL)
	grpcApp := grpcapp.New(log, grpcPort, authService)

	return &App{
		GRPCSrv: grpcApp,
	}
}
