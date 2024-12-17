package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del TaskService
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(title string) error {
	args := m.Called(title)
	return args.Error(0)
}

func (m *MockTaskService) GetAllTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func TestCreateTaskHandler(t *testing.T) {

	mockService := new(MockTaskService)
	mockService.On("CreateTask", "Test Task").Return(nil)

	handler := NewTaskHandler(mockService)

	payload := []byte(`{"title": "Test Task"}`)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.CreateTaskHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.JSONEq(t, `{"message": "Task created successfully"}`, rr.Body.String())
	mockService.AssertExpectations(t)
}

func TestCreateTaskHandler_Failure(t *testing.T) {

	mockService := new(MockTaskService)
	mockService.On("CreateTask", "Test Task").Return(errors.New("database error"))

	handler := NewTaskHandler(mockService)

	payload := []byte(`{"title": "Test Task"}`)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.CreateTaskHandler(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Failed to create task\n", rr.Body.String())
	mockService.AssertExpectations(t)
}

func TestGetAllTasksHandler(testing *testing.T) {

	mockService := new(MockTaskService)
	tasks := []models.Task{
		{ID: 1, Title: "Test Task 1", Done: false},
		{ID: 2, Title: "Test Task 2", Done: true},
	}
	mockService.On("GetAllTasks").Return(tasks, nil)

	handler := NewTaskHandler(mockService)

	req, err := http.NewRequest("GET", "/tasks", nil)
	assert.NoError(testing, err)

	rr := httptest.NewRecorder()

	handler.GetAllTasksHandler(rr, req)

	assert.Equal(testing, http.StatusOK, rr.Code)

	expectedResponse, _ := json.Marshal(tasks)
	assert.JSONEq(testing, string(expectedResponse), rr.Body.String())
	mockService.AssertExpectations(testing)
}
