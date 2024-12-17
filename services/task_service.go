package services

import "todo-api/models"

type TaskService interface {
	CreateTask(title string) error
	GetAllTasks() ([]models.Task, error)
}
