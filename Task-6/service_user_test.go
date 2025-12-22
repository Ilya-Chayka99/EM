package Task_6

import (
	"testing"
)

type UserServiceMock struct{}

func (u UserServiceMock) GetUser(id string) string {
	return id
}

func TestGetUserById(t *testing.T) {
	mock := UserServiceMock{}
	user := mock.GetUser("8438-48347-sdfsdf-3r8u23")
	if user != "8438-48347-sdfsdf-3r8u23" {
		t.Fatalf("expected mock-user, got %s", user)
	}
}
