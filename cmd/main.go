package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sparhokm/go-course-ms-auth/internal/app"
	"github.com/sparhokm/go-course-ms-auth/internal/config"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	application, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	go func() {
		application.Run()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	fmt.Println("Gracefully stopped")
}
