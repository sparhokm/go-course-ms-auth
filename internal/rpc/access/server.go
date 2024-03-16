package access

import (
	"google.golang.org/grpc"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/access_v1"
)

type server struct {
	desc.UnimplementedAccessV1Server
	access AccessService
}

func NewImplementation(access AccessService) *server {
	return &server{access: access}
}

func Register(gRPCServer *grpc.Server, access AccessService) {
	desc.RegisterAccessV1Server(gRPCServer, NewImplementation(access))
}
