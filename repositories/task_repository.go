package repositories

import "todo-api/models"

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
}
