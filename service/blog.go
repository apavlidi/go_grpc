package service

import (
	"errors"
	"go-grpc/proto"
	"slices"
	"time"
)

var blogs = []*proto.Blog{
	{
		Id:     1,
		Title:  "First Blog Post",
		Text:   "This is the content of the first blog post.",
		Author: nil,
		Date:   "2023-11-01",
	},
}

func ListBlogs() ([]*proto.Blog, error) {
	return blogs, nil
}

func CreateBlog(title string, text string, author *proto.User) (int, error) {
	id := IdGenerator()
	currentDate := time.Now().Local()

	foundUser := slices.IndexFunc(users, func(b *proto.User) bool {
		return b.Name == author.Name
	})
	if foundUser == -1 {
		return -1, errors.New("user not found")
	}

	blogs = append(blogs, &proto.Blog{
		Id:     int32(id),
		Title:  title,
		Text:   text,
		Author: author,
		Date:   currentDate.Format("2006-01-02"),
	})

	return id, nil
}

func DeleteBlog(id int32) {
	foundIndex := slices.IndexFunc(blogs, func(b *proto.Blog) bool {
		return b.Id == id
	})

	if foundIndex == -1 {
		return
	}

	blogs = append(blogs[:foundIndex], blogs[foundIndex+1:]...)
}

func GetBlog(id int32) *proto.Blog {
	foundIndex := slices.IndexFunc(blogs, func(b *proto.Blog) bool {
		return b.Id == id
	})

	if foundIndex == -1 {
		return nil
	}

	return blogs[foundIndex]
}

func UpdateBlog(req *proto.UpdateBlogRequest) {
	blog := GetBlog(req.Id)

	if blog != nil {
		blog.Title = req.Title
		blog.Author = req.Author
		blog.Text = req.Text
	}
}
