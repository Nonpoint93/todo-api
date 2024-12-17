// services/task_service.go

package services

import (
	"todo-api/models"
	"todo-api/repositories"
)

// TaskService maneja la lógica de negocio para las tareas.
type TaskService struct {
	Repo repositories.TaskRepository
}

// NewTaskService crea una instancia de TaskService.
func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) CreateTask(title string) error {
	task := &models.Task{Title: title, Done: false}
	return s.Repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.Repo.GetAllTasks()
}
