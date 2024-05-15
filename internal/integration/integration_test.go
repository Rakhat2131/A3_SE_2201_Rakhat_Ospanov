package integration

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "io/ioutil"
)

func TestRegisterAndCreateProject(t *testing.T) {
    req := httptest.NewRequest("POST", "/register", strings.NewReader(`{"username":"newuser","password":"password123"}`))
    w := httptest.NewRecorder()
    RegisterUserHandler(w, req)
    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }

    req = httptest.NewRequest("POST", "/projects", strings.NewReader(`{"name":"Integrated Project","description":"Project Description"}`))
    w = httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")
    CreateProjectHandler(w, req)
    resp = w.Result()
    body, _ = ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }
}

func TestLoginAndGetProjects(t *testing.T) {
    req := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"existinguser","password":"password123"}`))
    w := httptest.NewRecorder()
    LoginUserHandler(w, req)
    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }

    var token string
    // Extract token from response (omitted for brevity)

    req = httptest.NewRequest("GET", "/projects", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    w = httptest.NewRecorder()
    GetProjectsHandler(w, req)
    resp = w.Result()
    body, _ = ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }
}

func TestCreateProjectAndAddTask(t *testing.T) {
    req := httptest.NewRequest("POST", "/projects", strings.NewReader(`{"name":"Project for Tasks","description":"Description"}`))
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")
    CreateProjectHandler(w, req)
    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }

    var projectID int
    // Extract projectID from response (omitted for brevity)

    req = httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name":"Task for Project","description":"Task Description","projectID":projectID}`))
    w = httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")
    CreateTaskHandler(w, req)
    resp = w.Result()
    body, _ = ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }
}

func TestCreateProjectAndGetTasks(t *testing.T) {
    req := httptest.NewRequest("POST", "/projects", strings.NewReader(`{"name":"Old Project","description":"Description"}`))
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")
    CreateProjectHandler(w, req)
    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }

    var projectID int
    // Extract projectID from response (omitted for brevity)

    req = httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name":"Task for Old Project","description":"Task Description","projectID":projectID}`))
    w = httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")
    CreateTaskHandler(w, req)
    resp = w.Result()
    body, _ = ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }

    req = httptest.NewRequest("GET", "/projects/"+strconv.Itoa(projectID)+"/tasks", nil)
    w = httptest.NewRecorder()
    GetTasksHandler(w, req)
    resp = w.Result()
    body, _ = ioutil.ReadAll(resp.Body)
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
    }

    var tasks []Task
    // Unmarshal JSON response to get tasks (omitted for brevity)
    if len(tasks) != 1 || tasks[0].Name != "Task for Old Project" {
        t.Fatalf("Expected task name 'Task for Old Project', got %v", tasks)
    }
}
