package auth

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, userName string, password string) (string, error)
	GetRefreshToken(ctx context.Context, oldToken string) (string, error)
	GetAccessToken(ctx context.Context, token string) (string, error)
}
