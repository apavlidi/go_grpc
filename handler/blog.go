package handler

import (
	"context"
	"go-grpc/proto"
	"go-grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type BlogServer struct {
	proto.UnimplementedBlogServiceServer
}

func (s BlogServer) ListBlogs(_ context.Context, _ *proto.ListBlogRequest) (*proto.ListBlogResponse, error) {
	blogs, err := service.List()
	if err != nil {
		log.Fatalf("cannot get blogs %s", err)
	}
	return &proto.ListBlogResponse{Blogs: blogs}, nil
}

func (s BlogServer) CreateBlog(_ context.Context, req *proto.CreateBlogRequest) (*proto.CreateBlogResponse, error) {
	if req.Title == "" || req.Text == "" || req.Author == "" {
		return nil, status.New(codes.InvalidArgument, "invalid request").Err()
	}
	newBlogId := service.Create(req.Title, req.Text, req.Author)
	return &proto.CreateBlogResponse{Id: int32(newBlogId)}, nil
}

func (s BlogServer) DeleteBlog(_ context.Context, req *proto.DeleteBlogRequest) (*proto.DeleteBlogResponse, error) {
	service.Delete(req.Id)
	return &proto.DeleteBlogResponse{}, nil
}
