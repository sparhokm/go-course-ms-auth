package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func (s *server) Delete(ctx context.Context, req *desc.DeleteIn) (*emptypb.Empty, error) {
	err := s.user.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
