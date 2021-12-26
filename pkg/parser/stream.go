package parser

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"parser/pkg/models"
	"parser/pkg/repositories"
)

type Entry struct {
	Error error
	User  models.User
}

type Stream struct {
	Entry chan Entry
}

func Init() Stream {
	return Stream{
		Entry: make(chan Entry),
	}
}

func (s Stream) Watch() <-chan Entry {
	return s.Entry
}

func (s Stream) Start(path string, DB *sql.DB) {
	defer close(s.Entry)

	file, err := os.Open(path)

	if err != nil {
		s.Entry <- Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening delimiter. `[` or `{`
	if _, err := decoder.Token(); err != nil {
		s.Entry <- Entry{Error: fmt.Errorf("decode opening delimiter: %w", err)}
		return
	}

	i := 1

	for decoder.More() {
		var user models.User

		if err := decoder.Decode(&user); err != nil {
			s.Entry <- Entry{Error: fmt.Errorf("decode line %d: %w", i, err)}
			return
		}

		s.Entry <- Entry{User: user}

		repository := repositories.UserRepository{DB: DB}

		repository.CreateOne(&user)

		i++
	}

	// Read closing delimiter. `]` or `}`
	if _, err := decoder.Token(); err != nil {
		s.Entry <- Entry{Error: fmt.Errorf("decode closing delimiter: %w", err)}
		return
	}
}
