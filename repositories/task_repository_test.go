package repositories

import (
	"regexp"
	"testing"
	"todo-api/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(testing *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		testing.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer db.Close()

	repo := NewTaskRepositoryImpl(db)

	task := &models.Task{
		Title: "Test Task",
		Done:  false,
	}

	mock.ExpectExec(regexp.QuoteMeta(createTaskSQL)).
		WithArgs(task.Title, task.Done).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateTask(task)

	assert.NoError(testing, mock.ExpectationsWereMet())
}

func TestGetAllTasks(testing *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		testing.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer db.Close()

	repo := NewTaskRepositoryImpl(db)

	rows := sqlmock.NewRows([]string{"id", "title", "done"}).
		AddRow(1, "Task 1", false).
		AddRow(2, "Task 2", true)

	mock.ExpectQuery(regexp.QuoteMeta(getAllTasksSQL)).
		WillReturnRows(rows)

	tasks, err := repo.GetAllTasks()

	assert.NoError(testing, err)
	assert.Equal(testing, 2, len(tasks))
}
