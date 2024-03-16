package auth

import (
	"context"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/auth_v1"
)

func (s *server) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenIn) (*desc.GetRefreshTokenOut, error) {
	token, err := s.auth.GetRefreshToken(ctx, req.GetRefreshToken())

	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenOut{RefreshToken: token}, nil

}
