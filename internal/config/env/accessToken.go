package env

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	accessTokenTtl    = "ACCESS_TOKEN_TTL"
	accessTokenSecret = "ACCESS_TOKEN_SECRET"
)

type accessTokenConfig struct {
	secret       string
	timeDuration time.Duration
}

func NewAccessTokenConfig() (*accessTokenConfig, error) {
	accessTokenTTL, err := getTokenTtl(accessTokenTtl)
	if err != nil {
		return nil, fmt.Errorf("access token: %w", err)
	}

	secret := os.Getenv(accessTokenSecret)
	if len(secret) == 0 {
		return nil, errors.New("access token secret not found")
	}

	return &accessTokenConfig{
		timeDuration: *accessTokenTTL,
		secret:       secret,
	}, nil
}

func (cfg *accessTokenConfig) GetTimeDuration() time.Duration {
	return cfg.timeDuration
}

func (cfg *refreshTokenConfig) GetSecret() []byte {
	return []byte(cfg.secret)
}
