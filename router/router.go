package router

import (
	"net/http"
	"todo-api/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(handler *handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", handler.CreateTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/tasks", handler.GetAllTasksHandler).Methods(http.MethodGet)

	return r
}
