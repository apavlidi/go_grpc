package service

import (
	"go-grpc/proto"
	"time"
)

var blogs []*proto.Blog

func List() ([]*proto.Blog, error) {
	return []*proto.Blog{
		{
			Id:     1,
			Title:  "First Blog Post",
			Text:   "This is the content of the first blog post.",
			Author: "Author A",
			Date:   "2023-11-01",
		},
	}, nil
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
