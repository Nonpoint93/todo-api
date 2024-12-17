// handlers/task_handler.go

package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/services"
)

// TaskHandler define las dependencias necesarias.
type TaskHandler struct {
	Service *services.TaskService
}

// NewTaskHandler crea una instancia de TaskHandler.
func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateTask(input.Title); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Task created successfully"}`))
}

func (h *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}
