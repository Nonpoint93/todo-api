package main

import (
	"log"
	"net/http"
	"todo-api/config"
	"todo-api/handlers"
	"todo-api/repositories"
	"todo-api/router"
	"todo-api/services"
)

func main() {

	dbConfig, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	taskRepo := repositories.NewTaskRepositoryImpl(dbConfig.DB)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := router.NewRouter(taskHandler)

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
