package tracing

import (
	"github.com/uber/jaeger-client-go/config"
)

const appName = "ms_auth"

func Init() error {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(appName)

	return err
}
