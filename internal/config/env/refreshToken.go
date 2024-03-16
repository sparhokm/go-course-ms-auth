package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	refreshTokenTtl    = "REFRESH_TOKEN_TTL"
	refreshTokenSecret = "REFRESH_TOKEN_SECRET"
)

type refreshTokenConfig struct {
	secret       string
	timeDuration time.Duration
}

func NewRefreshTokenConfig() (*refreshTokenConfig, error) {
	refreshTokenTTL, err := getTokenTtl(refreshTokenTtl)
	if err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}

	secret := os.Getenv(refreshTokenSecret)
	if len(secret) == 0 {
		return nil, errors.New("refresh token secret not found")
	}

	return &refreshTokenConfig{
		timeDuration: *refreshTokenTTL,
		secret:       secret,
	}, nil
}

func (cfg *refreshTokenConfig) GetTimeDuration() time.Duration {
	return cfg.timeDuration
}

func (cfg *accessTokenConfig) GetSecret() []byte {
	return []byte(cfg.secret)
}

func getTokenTtl(key string) (*time.Duration, error) {
	tokenTtl := os.Getenv(key)
	if len(tokenTtl) == 0 {
		return nil, errors.New("grpc host not found")
	}
	ttl, err := strconv.Atoi(tokenTtl)
	if err != nil {
		return nil, err
	}

	timeDuration := time.Duration(ttl) * time.Second

	return &timeDuration, err
}
