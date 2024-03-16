package auth

import (
	"google.golang.org/grpc"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/auth_v1"
)

type server struct {
	desc.UnimplementedAuthV1Server
	auth AuthService
}

func NewImplementation(auth AuthService) *server {
	return &server{auth: auth}
}

func Register(gRPCServer *grpc.Server, auth AuthService) {
	desc.RegisterAuthV1Server(gRPCServer, NewImplementation(auth))
}
