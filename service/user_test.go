package service

import (
	"testing"
)

func TestCreateUserValidName(t *testing.T) {
	beforeCount := len(users)
	id := CreateUser("Alice")
	afterCount := len(users)

	if id == 0 {
		t.Errorf("expected valid ID, got 0")
	}

	if afterCount != beforeCount+1 {
		t.Errorf("expected user list size to increase, but found %d and %d", beforeCount, afterCount)
	}

	lastUser := users[len(users)-1]
	if lastUser.Name != "Alice" {
		t.Errorf("expected user name Alice, got %s", lastUser.Name)
	}
}
