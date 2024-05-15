package users

import (
    "testing"
    "reflect"
)

func TestCreateUserSuccess(t *testing.T) {
    user := User{Username: "newuser", Password: "password123"}
    expected := User{ID: 1, Username: "newuser", Password: "hashedpassword123"}

    // Assuming CreateUser returns a User object and nil error on success
    result, err := CreateUser(user)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestCreateUserDuplicate(t *testing.T) {
    user := User{Username: "existinguser", Password: "password123"}
    
    // Assuming CreateUser returns an error on duplicate username
    _, err := CreateUser(user)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}
