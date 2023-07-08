package repositories

import (
	"s3-microservice/models"
)

type UserRepository struct {
	users []*models.User
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	mockUsers := []*models.User{
		{ID: 1, FirstName: "Alice", LastName: "Smith"},
		{ID: 2, FirstName: "Bob", LastName: "Jones"},
		{ID: 3, FirstName: "Charlie", LastName: "Thompson"},
	}
	return mockUsers, nil
}
