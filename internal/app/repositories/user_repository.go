package repositories

import (
	"database/sql"
	"fmt"
	"parser/internal/app/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) CreateOne(user *models.User) {
	sqlStatement := `INSERT INTO users (username, name, email) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := repository.DB.QueryRow(sqlStatement, user.Username, user.Name, user.Email).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("User %d created", id))
}
