package auth

import (
	"context"
	"errors"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/auth_v1"
)

func (s *server) Login(ctx context.Context, req *desc.LoginIn) (*desc.LoginOut, error) {
	if req.GetPassword() == "" || req.GetEmail() == "" {
		return nil, errors.New("wrong input data")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &desc.LoginOut{RefreshToken: token}, nil
}
