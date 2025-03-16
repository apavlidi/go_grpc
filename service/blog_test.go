package service

import (
	"github.com/stretchr/testify/assert"
	"go-grpc/proto"
	"testing"
)

func Test_ListBlogs(t *testing.T) {
	listBlogs, err := ListBlogs()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(listBlogs))
}

func Test_CreateBlog(t *testing.T) {
	listBlogs, err := ListBlogs()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(listBlogs))

	_ = CreateUser("apavlidi")
	_, err = CreateBlog("a title", "a text", &proto.User{Name: "apavlidi"})
	assert.Nil(t, err)

	updatedListBlogs, err := ListBlogs()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(updatedListBlogs))
}
