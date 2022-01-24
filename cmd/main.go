package main

import (
	"github.com/joho/godotenv"
	"log"
	"parser/config"
	"parser/internal/app/db"
	"parser/internal/app/models"
	"parser/internal/app/reader"
	"parser/internal/app/repositories"
	"parser/internal/app/services"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.NewConfig()
	postgres := db.Postgres{DbConfig: conf.DbConfig}

	db, err := postgres.Connect()

	defer db.Close()

	if err != nil {
		panic(err)
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	fileReader := reader.NewFileReader("users.json")

	dataCh := make(chan models.User)

	go fileReader.Read(dataCh)

	userService.Handle(dataCh)
}
