package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetIn) (*desc.GetOut, error) {
	_ = ctx
	log.Printf("Get user id: %d", req.GetId())

	return &desc.GetOut{
		Id: req.GetId(),
		UserInfo: &desc.UserIno{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Role:  desc.Role(1),
		},
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateIn) (*desc.CreateOut, error) {
	_ = ctx
	log.Printf("Create user %+v", req)

	return &desc.CreateOut{Id: 1}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateIn) (*emptypb.Empty, error) {
	_ = ctx
	log.Printf("Update user %+v", req)

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteIn) (*emptypb.Empty, error) {
	_ = ctx
	log.Printf("Delete user id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
