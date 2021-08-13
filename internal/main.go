package main

import (
	"github.com/joho/godotenv"
	todoapi "github.com/klaus-abram/todo-rest-api"
	"github.com/klaus-abram/todo-rest-api/pkg/handler"
	"github.com/klaus-abram/todo-rest-api/pkg/repository"
	"github.com/klaus-abram/todo-rest-api/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialising configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables %s", err.Error())
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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	serv := new(todoapi.Server)
	if errRun := serv.RunServer(viper.GetString("port"), handlers.InitRoutes()); errRun != nil {
		log.Fatalf("error with the running the http server %s", errRun.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
