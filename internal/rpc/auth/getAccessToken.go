package auth

import (
	"context"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/auth_v1"
)

func (s *server) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenIn) (*desc.GetAccessTokenOut, error) {
	token, err := s.auth.GetAccessToken(ctx, req.GetRefreshToken())

	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenOut{AccessToken: token}, nil

}
