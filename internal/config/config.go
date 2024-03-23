package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/sparhokm/go-course-ms-auth/internal/config/env"
)

type Config struct {
	GRPCConfig         GRPCConfig
	PGConfig           PGConfig
	AccessTokenConfig  TokenConfig
	PrometheusConfig   PrometheusConfig
	RefreshTokenConfig TokenConfig
}

func MustLoad() *Config {
	path := fetchConfigPath()
	err := godotenv.Load(path)

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	prometheusConfig, err := env.NewPrometheusConfig()
	if err != nil {
		log.Fatalf("failed to get prometheus config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	accessTokenConfig, err := env.NewAccessTokenConfig()
	if err != nil {
		log.Fatalf("failed to get access token config: %v", err)
	}

	refreshTokenConfig, err := env.NewRefreshTokenConfig()
	if err != nil {
		log.Fatalf("failed to get access token config: %v", err)
	}

	return &Config{
		GRPCConfig:         grpcConfig,
		PrometheusConfig:   prometheusConfig,
		PGConfig:           pgConfig,
		AccessTokenConfig:  accessTokenConfig,
		RefreshTokenConfig: refreshTokenConfig,
	}
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", ".env", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
