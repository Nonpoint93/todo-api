package router

import (
	"net/http"
	"todo-api/config"
	"todo-api/handlers"
	"todo-api/repositories"
	"todo-api/services"

	"github.com/gorilla/mux"
)

func NewRouter(conf *config.Config, dbConfig *config.DatabaseConfig) *mux.Router {

	taskRepo := repositories.NewTaskRepositoryImpl(dbConfig.DB)
	taskService := services.NewTaskServiceImpl(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := mux.NewRouter()

	for _, endpoint := range conf.Endpoints {

		switch endpoint.Handler {
		case "GetTasks":
			r.HandleFunc(endpoint.Path, taskHandler.GetAllTasksHandler).Methods(endpoint.Method)
		case "CreateTask":
			r.HandleFunc(endpoint.Path, taskHandler.CreateTaskHandler).Methods(endpoint.Method)
		default:
			r.HandleFunc(endpoint.Path, func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Handler not implemented", http.StatusNotImplemented)
			}).Methods(endpoint.Method)
		}
	}
	return r
}
