package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"parsing/pkg/config"
	"parsing/pkg/db"
	"parsing/pkg/models"
	"parsing/pkg/parser"
	"parsing/pkg/repositories"
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

	var users models.Users

	json.Unmarshal([]byte(data), &users)

	repositories.CreateUsers(users, db)
}
