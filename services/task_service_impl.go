package services

import (
	"todo-api/models"
	"todo-api/repositories"
)

type TaskServiceImpl struct {
	Repo repositories.TaskRepository
}

func NewTaskServiceImpl(repo repositories.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{Repo: repo}
}

func (taskServiceImpl *TaskServiceImpl) CreateTask(title string) error {
	task := &models.Task{Title: title, Done: false}
	return taskServiceImpl.Repo.CreateTask(task)
}

func (taskServiceImpl *TaskServiceImpl) GetAllTasks() ([]models.Task, error) {
	return taskServiceImpl.Repo.GetAllTasks()
}
