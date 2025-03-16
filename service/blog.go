package service

import (
	"go-grpc/proto"
	"slices"
	"time"
)

var blogs = []*proto.Blog{
	{
		Id:     1,
		Title:  "First Blog Post",
		Text:   "This is the content of the first blog post.",
		Author: "Author A",
		Date:   "2023-11-01",
	},
}

func List() ([]*proto.Blog, error) {
	return blogs, nil
}

func Create(title string, text string, author string) int {
	id := IdGenerator()
	currentDate := time.Now().Local()

	blogs = append(blogs, &proto.Blog{
		Id:     int32(id),
		Title:  title,
		Text:   text,
		Author: author,
		Date:   currentDate.Format("2006-01-02"),
	})

	return id
}

func Delete(id int32) {
	foundIndex := slices.IndexFunc(blogs, func(b *proto.Blog) bool {
		return b.Id == id
	})

	if foundIndex == -1 {
		return
	}

	blogs = append(blogs[:foundIndex], blogs[foundIndex+1:]...)
}
