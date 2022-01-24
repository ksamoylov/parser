package services

import (
	"fmt"
	"parser/internal/app/models"
	"parser/internal/app/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) Handle(dataCh <-chan models.User) {
	for {
		select {
		case user, ok := <-dataCh:
			if !ok {
				dataCh = nil
				break
			}

			userService.userRepository.CreateOne(&user)
		}

		if dataCh == nil {
			fmt.Println("Parsing completed")

			break
		}
	}
}
