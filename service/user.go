package service

import (
	"go-grpc/proto"
)

var Users = []*proto.User{}

func ListUsers() []*proto.User {
	return Users
}

func CreateUser(name string) int {
	id := IdGenerator()

	Users = append(Users, &proto.User{
		Id:   int32(id),
		Name: name,
	})

	return id
}
