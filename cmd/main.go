package main

//command execute db:
// docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres
// migrate -path ./schema -database postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable up

import (
	"context"
	"github.com/joho/godotenv"
	todoapi "github.com/klaus-abram/todo-rest-api"
	"github.com/klaus-abram/todo-rest-api/internal/handler"
	"github.com/klaus-abram/todo-rest-api/internal/repository"
	"github.com/klaus-abram/todo-rest-api/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Todo Rest API
// @version 1.0
// @description API Documentation for the rest-api Todo App

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialising configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	serv := new(todoapi.Server)
	go func() {
		if errRun := serv.RunServer(viper.GetString("port"), handlers.InitRoutes()); errRun != nil {
			logrus.Fatalf("error with the running the http server %s", errRun.Error())
		}
	}()

	logrus.Print("todo-rest-app - started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("todo-rest-app - shutting down")

	if err := serv.ShutdownServer(context.Background()); err != nil {
		logrus.Errorf("error with server shutting down %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db disconnect %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
