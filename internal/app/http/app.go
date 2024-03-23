package http

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/sparhokm/go-course-ms-auth/internal/config"
)

type App struct {
	httpServer *http.Server
}

func New(config config.PrometheusConfig) *App {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	prometheusServer := &http.Server{
		Addr:              config.Address(),
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	return &App{httpServer: prometheusServer}
}

func (a *App) MustRun() {
	err := a.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	// nolint: errcheck
	a.httpServer.Shutdown(context.Background())
}
