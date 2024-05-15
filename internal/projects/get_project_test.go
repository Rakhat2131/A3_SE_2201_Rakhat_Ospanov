package projects

import (
    "testing"
    "reflect"
)

func TestGetProjectSuccess(t *testing.T) {
    projectID := 1
    expected := Project{ID: 1, Name: "Existing Project", Description: "Description"}

    result, err := GetProject(projectID)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestGetProjectNotFound(t *testing.T) {
    projectID := 999

    _, err := GetProject(projectID)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}
