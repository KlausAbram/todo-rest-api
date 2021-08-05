package main

import (
	todoapi "github.com/klaus-abram/todo-rest-api"
	"github.com/klaus-abram/todo-rest-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	serv := new(todoapi.Server)
	if errRun := serv.RunServer("8000", handlers.InitRoutes()); errRun != nil {
		log.Fatalf("error with the running the http server %s", errRun.Error())
	}
}
