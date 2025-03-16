package handler

import (
	"context"
	"go-grpc/proto"
	"go-grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	proto.UnimplementedUserServiceServer
}

func (s UserHandler) ListUsers(_ context.Context, _ *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	users := service.ListUsers()
	return &proto.ListUsersResponse{Users: users}, nil
}

func (s UserHandler) CreateUser(_ context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	if req.Name == "" {
		return nil, status.New(codes.InvalidArgument, "invalid request").Err()
	}
	service.CreateUser(req.Name)
	return &proto.CreateUserResponse{}, nil
}
