package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func (s *server) Update(ctx context.Context, req *desc.UpdateIn) (*emptypb.Empty, error) {
	var (
		name, email *string
		role        *user.Role
		err         error
	)

	if req.GetName() != nil {
		name = &req.GetName().Value
	}

	if req.GetEmail() != nil {
		email = &req.GetEmail().Value
	}

	if req.Role != nil {
		role, err = user.NewRole(req.GetRole().String())
		if err != nil {
			return nil, err
		}
	}

	err = s.user.Update(ctx, req.GetId(), name, email, role)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
