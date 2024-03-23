package config

import "time"

type GRPCConfig interface {
	Address() string
}

type PrometheusConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type TokenConfig interface {
	GetSecret() []byte
	GetTimeDuration() time.Duration
}
