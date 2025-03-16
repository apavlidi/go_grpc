package handler

import (
	"context"
	"go-grpc/proto"
	"go-grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type BlogHandler struct {
	proto.UnimplementedBlogServiceServer
}

func (s BlogHandler) ListBlogs(_ context.Context, _ *proto.ListBlogsRequest) (*proto.ListBlogsResponse, error) {
	blogs, err := service.ListBlogs()
	if err != nil {
		log.Fatalf("cannot get blogs %s", err)
	}
	return &proto.ListBlogsResponse{Blogs: blogs}, nil
}

func (s BlogHandler) CreateBlog(_ context.Context, req *proto.CreateBlogRequest) (*proto.CreateBlogResponse, error) {
	if req.Title == "" || req.Text == "" || req.Author == "" {
		return nil, status.New(codes.InvalidArgument, "invalid request").Err()
	}
	newBlogId := service.CreateBlog(req.Title, req.Text, req.Author)
	return &proto.CreateBlogResponse{Id: int32(newBlogId)}, nil
}

func (s BlogHandler) DeleteBlog(_ context.Context, req *proto.DeleteBlogRequest) (*proto.DeleteBlogResponse, error) {
	service.DeleteBlog(req.Id)
	return &proto.DeleteBlogResponse{}, nil
}

func (s BlogHandler) GetBlog(_ context.Context, req *proto.GetBlogRequest) (*proto.GetBlogResponse, error) {
	blog := service.GetBlog(req.Id)
	return &proto.GetBlogResponse{Blog: blog}, nil
}

func (s BlogHandler) UpdateBlog(_ context.Context, req *proto.UpdateBlogRequest) (*proto.UpdateBlogResponse, error) {
	service.UpdateBlog(req)
	return &proto.UpdateBlogResponse{}, nil
}
