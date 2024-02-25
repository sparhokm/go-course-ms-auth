package user

import (
	"context"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func (s *server) Get(ctx context.Context, req *desc.GetIn) (*desc.GetOut, error) {
	user, err := s.user.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return ToUserRpc(user), nil
}
