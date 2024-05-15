package tasks

import (
    "testing"
    "reflect"
)

func TestCreateTaskSuccess(t *testing.T) {
    task := Task{Name: "New Task", Description: "Task Description", ProjectID: 1}
    expected := Task{ID: 1, Name: "New Task", Description: "Task Description", ProjectID: 1}

    result, err := CreateTask(task)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestCreateTaskInvalidProject(t *testing.T) {
    task := Task{Name: "New Task", Description: "Task Description", ProjectID: 999}

    _, err := CreateTask(task)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func CreateTask(task invalid type) {
	panic("unimplemented")
}
