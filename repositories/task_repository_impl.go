package repositories

import (
	"database/sql"
	"todo-api/models"
)

const (
	createTaskSQL  = `INSERT INTO tasks (title, done) VALUES (?, ?)`
	getAllTasksSQL = `SELECT id, title, done FROM tasks`
)

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepositoryImpl(db *sql.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) CreateTask(task *models.Task) error {
	_, err := r.db.Exec(createTaskSQL, task.Title, task.Done)
	return err
}

func (r *TaskRepositoryImpl) GetAllTasks() ([]models.Task, error) {
	rows, err := r.db.Query(getAllTasksSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
