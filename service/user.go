package service

import (
	"go-grpc/proto"
)

var users = []*proto.User{}

func ListUsers() []*proto.User {
	return users
}

func CreateUser(name string) int {
	id := IdGenerator()

	users = append(users, &proto.User{
		Id:   int32(id),
		Name: name,
	})

	return id
}
