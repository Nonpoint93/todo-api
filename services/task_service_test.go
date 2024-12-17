package services

import (
	"errors"
	"testing"
	"todo-api/mocks"
	"todo-api/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask_Success(testing *testing.T) {

	ctrl := gomock.NewController(testing)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	service := NewTaskServiceImpl(mockRepo)

	title := "New Task"
	expectedTask := &models.Task{Title: title, Done: false}

	mockRepo.EXPECT().CreateTask(expectedTask).Return(nil)

	err := service.CreateTask(title)

	assert.NoError(testing, err)
}

func TestCreateTask_Failure(testing *testing.T) {

	ctrl := gomock.NewController(testing)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	service := NewTaskServiceImpl(mockRepo)

	title := "New Task"
	expectedTask := &models.Task{Title: title, Done: false}

	mockRepo.EXPECT().CreateTask(expectedTask).Return(errors.New("database error"))

	err := service.CreateTask(title)

	assert.Error(testing, err)
	assert.Equal(testing, "database error", err.Error())
}

func TestGetAllTasks_Success(testing *testing.T) {

	ctrl := gomock.NewController(testing)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	service := NewTaskServiceImpl(mockRepo)

	expectedTasks := []models.Task{
		{ID: 1, Title: "Task 1", Done: false},
		{ID: 2, Title: "Task 2", Done: true},
	}

	mockRepo.EXPECT().GetAllTasks().Return(expectedTasks, nil)

	tasks, err := service.GetAllTasks()

	assert.NoError(testing, err)
	assert.Equal(testing, expectedTasks, tasks)
}

func TestGetAllTasks_Failure(testing *testing.T) {

	ctrl := gomock.NewController(testing)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	service := NewTaskServiceImpl(mockRepo)

	mockRepo.EXPECT().GetAllTasks().Return(nil, errors.New("database error"))

	tasks, err := service.GetAllTasks()

	assert.Error(testing, err)
	assert.Nil(testing, tasks)
	assert.Equal(testing, "database error", err.Error())
}
