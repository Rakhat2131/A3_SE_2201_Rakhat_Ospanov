package tasks

import (
    "testing"
    "reflect"
)

func TestGetTasksSuccess(t *testing.T) {
    projectID := 1
    expected := []Task{{ID: 1, Name: "Existing Task", Description: "Description", ProjectID: 1}}

    result, err := GetTasks(projectID)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestGetTasksInvalidProject(t *testing.T) {
    projectID := 999

    _, err := GetTasks(projectID)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func GetTasks(projectID int) {
	panic("unimplemented")
}
