package grpcapp

import (
	"fmt"
	authgrpc "github.com/malinatrash/grpc-auth-server/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int, authService authgrpc.Auth) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer, authService)
	return &App{
		log:        log,
		port:       port,
		gRPCServer: gRPCServer,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op))
	log.Info("starting gRPC server", slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"
	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
