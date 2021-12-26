package main

import (
	"github.com/joho/godotenv"
	"log"
	"parser/pkg/config"
	"parser/pkg/db"
	"parser/pkg/parser"
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

	stream := parser.Init()

	go func() {
		for data := range stream.Watch() {

			if data.Error != nil {
				panic(data.Error)
			}

			log.Println(data.User.ID, ":", data.User.Name)
		}
	}()

	stream.Start("users.json", db)
}
