package user

import (
	"context"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func (s *server) Create(ctx context.Context, req *desc.CreateIn) (*desc.CreateOut, error) {
	password, err := user.NewPassword(req.GetPassword().GetPassword(), req.GetPassword().GetConfirm())
	if err != nil {
		return nil, err
	}

	role, err := user.NewRole(req.GetUserInfo().GetRole().String())
	if err != nil {
		return nil, err
	}

	userInfo, err := user.NewUserInfo(req.GetUserInfo().GetName(), req.GetUserInfo().GetEmail(), *role)
	if err != nil {
		return nil, err
	}

	id, err := s.user.Create(ctx, userInfo, password)
	if err != nil {
		return nil, err
	}

	return &desc.CreateOut{Id: id}, nil
}
