package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	grpcapp "github.com/sparhokm/go-course-ms-auth/internal/app/grpc"
	"github.com/sparhokm/go-course-ms-auth/internal/config"
	accessRepo "github.com/sparhokm/go-course-ms-auth/internal/repository/access"
	"github.com/sparhokm/go-course-ms-auth/internal/repository/deletedUser"
	userRepo "github.com/sparhokm/go-course-ms-auth/internal/repository/user"
	"github.com/sparhokm/go-course-ms-auth/internal/service/access"
	"github.com/sparhokm/go-course-ms-auth/internal/service/auth"
	"github.com/sparhokm/go-course-ms-auth/internal/service/hasher"
	"github.com/sparhokm/go-course-ms-auth/internal/service/tokenGenerator"
	"github.com/sparhokm/go-course-ms-auth/internal/service/user"
	"github.com/sparhokm/go-course-ms-auth/pkg/client/db/pg"
	"github.com/sparhokm/go-course-ms-auth/pkg/client/db/transaction"
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
	aRepo := accessRepo.NewRepository(dbClient)
	deletedUserRepo := deletedUser.NewRepository(dbClient)
	txManager := transaction.NewTransactionManager(dbClient.DB())

	passwordHasher := hasher.NewHasher()
	accessTokenGenerator := tokenGenerator.NewTokenGenerator(config.AccessTokenConfig.GetSecret(), config.AccessTokenConfig.GetTimeDuration())
	refreshTokenGenerator := tokenGenerator.NewTokenGenerator(config.RefreshTokenConfig.GetSecret(), config.RefreshTokenConfig.GetTimeDuration())

	userService := user.NewUserService(uRepo, deletedUserRepo, passwordHasher, txManager)
	authService := auth.NewAuthService(uRepo, passwordHasher, accessTokenGenerator, refreshTokenGenerator)
	accessService := access.NewAccessService(accessTokenGenerator, aRepo)

	a.GRPCServer = grpcapp.New(config.GRPCConfig, userService, authService, accessService)
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
