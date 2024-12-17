package router

import (
	"net/http"
	"todo-api/handlers"

	"github.com/gorilla/mux"
)

const uriTask string = "/tasks"

func NewRouter(handler *handlers.TaskHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(uriTask, handler.CreateTaskHandler).Methods(http.MethodPost)
	router.HandleFunc(uriTask, handler.GetAllTasksHandler).Methods(http.MethodGet)

	return router
}
