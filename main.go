package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-api/config"
	router "todo-api/routers"
)

func main() {

	conf, err := config.LoadConfig("application.yml")
	if nil != err {
		log.Fatalf("Could not initialize config from application properties, %v", err)
	}

	dbConfig, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	r := router.NewRouter(conf, dbConfig)

	log.Printf("Server running on port %d", conf.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Server.Port), r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
