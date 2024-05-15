package projects

import (
    "testing"
    "reflect"
)

func TestCreateProjectSuccess(t *testing.T) {
    project := Project{Name: "New Project", Description: "Project Description"}
    expected := Project{ID: 1, Name: "New Project", Description: "Project Description"}

    result, err := CreateProject(project)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestCreateProjectMissingName(t *testing.T) {
    project := Project{Description: "Project Description"}

    _, err := CreateProject(project)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}
