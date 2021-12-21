package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"parser/pkg/config"
	"parser/pkg/db"
	"parser/pkg/models"
	"parser/pkg/parser"
	"parser/pkg/repositories"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()

	postgres := db.Postgres{DbConfig: conf.DbConfig}

	db, err := postgres.Connect()

	defer db.Close()

	if err != nil {
		return
	}

	data := parser.Parse("users.json")

	var users []models.User

	json.Unmarshal([]byte(data), &users)

	userRepository := repositories.UserRepository{DB: db}

	userRepository.CreateMany(users)
}
