package grpcapp

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sparhokm/go-course-ms-auth/internal/config"
	interceptor "github.com/sparhokm/go-course-ms-auth/internal/incerceptor"
	"github.com/sparhokm/go-course-ms-auth/internal/rpc/access"
	"github.com/sparhokm/go-course-ms-auth/internal/rpc/auth"
	"github.com/sparhokm/go-course-ms-auth/internal/rpc/user"
)

type App struct {
	gRPCServer *grpc.Server
	config     config.GRPCConfig
}

func New(
	config config.GRPCConfig,
	userService user.UserService,
	authService auth.AuthService,
	accessService access.AccessService,
) *App {
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptor.LogInterceptor,
		interceptor.MetricsInterceptor,
		interceptor.ServerTracingInterceptor,
	))
	reflection.Register(gRPCServer)
	user.Register(gRPCServer, userService)
	auth.Register(gRPCServer, authService)
	access.Register(gRPCServer, accessService)

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
