package user

import (
	"google.golang.org/grpc"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

type server struct {
	desc.UnimplementedUserV1Server
	user UserService
}

func NewImplementation(user UserService) *server {
	return &server{user: user}
}

func Register(gRPCServer *grpc.Server, user UserService) {
	desc.RegisterUserV1Server(gRPCServer, NewImplementation(user))
}
