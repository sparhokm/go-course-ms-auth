package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	grpcapp "github.com/sparhokm/go-course-ms-auth/internal/app/grpc"
	"github.com/sparhokm/go-course-ms-auth/internal/client/db/pg"
	"github.com/sparhokm/go-course-ms-auth/internal/client/db/transaction"
	"github.com/sparhokm/go-course-ms-auth/internal/config"
	"github.com/sparhokm/go-course-ms-auth/internal/repository/deletedUser"
	userRepo "github.com/sparhokm/go-course-ms-auth/internal/repository/user"
	"github.com/sparhokm/go-course-ms-auth/internal/service"
	"github.com/sparhokm/go-course-ms-auth/internal/service/user"
)

type App struct {
	GRPCServer *grpcapp.App
	Dbc        *pgxpool.Pool
}

func NewApp(ctx context.Context, config *config.Config) (*App, error) {
	a := &App{}

	dbc, err := pgxpool.New(ctx, config.PGConfig.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	err = dbc.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping error: %s", err.Error())
	}

	dbClient := pg.New(dbc)

	uRepo := userRepo.NewRepository(dbClient)
	deletedUserRepo := deletedUser.NewRepository(dbClient)
	txManager := transaction.NewTransactionManager(dbClient.DB())

	passwordHasher := service.NewHasher()
	userService := user.NewUserService(uRepo, deletedUserRepo, passwordHasher, txManager)

	a.GRPCServer = grpcapp.New(config.GRPCConfig, userService)
	return a, nil
}

func (a App) Run() {
	go func() {
		a.GRPCServer.MustRun()
	}()
}

func (a App) Stop() {
	a.GRPCServer.Stop()
	a.Dbc.Close()
}
