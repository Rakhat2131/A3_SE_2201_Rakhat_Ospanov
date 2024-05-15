package users

import (
    "testing"
    "reflect"
)

func TestAuthenticateSuccess(t *testing.T) {
    user := User{Username: "testuser", Password: "correctpassword"}
    expected := "valid_token"

    token, err := Authenticate(user.Username, user.Password)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if token != expected {
        t.Errorf("Expected %v, got %v", expected, token)
    }
}

func TestAuthenticateFailure(t *testing.T) {
    user := User{Username: "testuser", Password: "wrongpassword"}

    _, err := Authenticate(user.Username, user.Password)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}
