package repositories

import (
	"database/sql"
	"fmt"
	"parsing/pkg/models"
)

func CreateUsers(users models.Users, db *sql.DB) {
	for _, user := range users {
		sqlStatement := `INSERT INTO users (username, name, email) VALUES ($1, $2, $3) RETURNING id`
		id := 0
		err := db.QueryRow(sqlStatement, user.Username, user.Name, user.Email).Scan(&id)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
	}
}
