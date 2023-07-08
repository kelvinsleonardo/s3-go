package services

import (
	models "s3-microservice/models"
	"s3-microservice/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func (us *UserService) GetAllUsers() ([]*models.User, error) {
	return us.userRepo.GetAllUsers()
}
