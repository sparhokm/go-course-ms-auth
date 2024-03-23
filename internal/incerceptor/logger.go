package interceptor

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"

	"github.com/sparhokm/go-course-ms-auth/internal/infra/logger"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), slog.String("method", info.FullMethod), slog.Any("req", req))
	}

	logger.Info("request", slog.String("method", info.FullMethod), slog.Any("req", req), slog.Any("res", res), slog.Duration("duration", time.Since(now)))

	return res, err
}
