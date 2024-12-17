package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/services"
)

type TaskHandler struct {
	Service services.TaskService
}

func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (taskHandler *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := taskHandler.Service.CreateTask(input.Title); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Task created successfully"}`))
}

func (taskHandler *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskHandler.Service.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}
