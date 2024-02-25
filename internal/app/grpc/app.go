package grpcapp

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sparhokm/go-course-ms-auth/internal/config"
	"github.com/sparhokm/go-course-ms-auth/internal/rpc/user"
)

type App struct {
	gRPCServer *grpc.Server
	config     config.GRPCConfig
}

func New(config config.GRPCConfig, userService user.UserService) *App {
	gRPCServer := grpc.NewServer()
	reflection.Register(gRPCServer)
	user.Register(gRPCServer, userService)

	return &App{
		gRPCServer: gRPCServer,
		config:     config,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", a.config.Address())
	if err != nil {
		return fmt.Errorf("grpcapp.Run: %w", err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("grpcapp.Run: %w", err)
	}

	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
