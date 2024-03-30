package access

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/access_v1"
)

const authPrefix = "Bearer "

func (s *server) Check(ctx context.Context, req *desc.CheckIn) (*desc.CheckOut, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	id, err := s.access.Check(ctx, accessToken, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &desc.CheckOut{Id: id}, nil
}
