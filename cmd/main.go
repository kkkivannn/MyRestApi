package main

import (
	api "RestaurantRestApi"
	"RestaurantRestApi/pkg/handler"
	"RestaurantRestApi/pkg/repository"
	"RestaurantRestApi/pkg/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err.Error())
	}

	db, err := repository.NewPostgresDb(&repository.ConfigDb{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("Не получилось инициализировтаь базу: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	//manager := auth.NewManager()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	var server = new(api.Server)
	if err := server.Run(os.Getenv("DB_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Server not running: %s", err.Error())
	}
}
