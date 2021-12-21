package repositories

import (
	"database/sql"
	"fmt"
	"parser/pkg/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (repository *UserRepository) CreateMany(users []models.User) {
	for _, user := range users {
		repository.createOne(&user)
	}
}

func (repository *UserRepository) createOne(user *models.User) {
	sqlStatement := `INSERT INTO users (username, name, email) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := repository.DB.QueryRow(sqlStatement, user.Username, user.Name, user.Email).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Id: ", id)
}
