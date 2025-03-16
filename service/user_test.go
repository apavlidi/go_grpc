package service

import (
	"testing"
)

func TestCreateUserValidName(t *testing.T) {
	beforeCount := len(Users)
	id := CreateUser("Alice")
	afterCount := len(Users)

	if id == 0 {
		t.Errorf("expected valid ID, got 0")
	}

	if afterCount != beforeCount+1 {
		t.Errorf("expected user list size to increase, but found %d and %d", beforeCount, afterCount)
	}

	lastUser := Users[len(Users)-1]
	if lastUser.Name != "Alice" {
		t.Errorf("expected user name Alice, got %s", lastUser.Name)
	}
}
