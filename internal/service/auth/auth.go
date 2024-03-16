package auth

import (
	"context"
)

type service struct {
	userRepo              UserRepo
	hasher                Hasher
	accessTokenGenerator  TokenGenerator
	refreshTokenGenerator TokenGenerator
}

func NewAuthService(userRepo UserRepo, hasher Hasher, accessTokenGenerator TokenGenerator, refreshTokenGenerator TokenGenerator) *service {
	return &service{
		userRepo:              userRepo,
		hasher:                hasher,
		accessTokenGenerator:  accessTokenGenerator,
		refreshTokenGenerator: refreshTokenGenerator,
	}
}

func (a service) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := a.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = a.hasher.Verify(user.PasswordHash, password)
	if err != nil {
		return "", err
	}

	return a.refreshTokenGenerator.GenerateToken(user.Id, user.Info.Role)
}

func (a service) GetRefreshToken(ctx context.Context, oldToken string) (string, error) {
	claims, err := a.refreshTokenGenerator.VerifyToken(oldToken)
	if err != nil {
		return "", err
	}

	u, err := a.userRepo.Get(ctx, claims.Id)
	if err != nil {
		return "", err
	}

	return a.refreshTokenGenerator.GenerateToken(u.Id, u.Info.Role)
}

func (a service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := a.refreshTokenGenerator.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	u, err := a.userRepo.Get(ctx, claims.Id)
	if err != nil {
		return "", err
	}

	return a.accessTokenGenerator.GenerateToken(u.Id, u.Info.Role)
}
