package main

import (
	todoapi "github.com/klaus-abram/todo-rest-api"
	"github.com/klaus-abram/todo-rest-api/pkg/handler"
	"github.com/klaus-abram/todo-rest-api/pkg/repository"
	"github.com/klaus-abram/todo-rest-api/pkg/service"
	"log"
)

func main() {

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	serv := new(todoapi.Server)
	if errRun := serv.RunServer("8000", handlers.InitRoutes()); errRun != nil {
		log.Fatalf("error with the running the http server %s", errRun.Error())
	}
}
