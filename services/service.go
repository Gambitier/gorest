package services

import (
	"github.com/gambitier/gorest/database"
	userservice "github.com/gambitier/gorest/services/users"
)

type Service struct {
	UserService *userservice.UserService
}

func NewService(db *database.DatabaseRepository) (*Service, error) {
	userService := userservice.NewUserService(db.UserRepo)
	return &Service{
		UserService: userService,
	}, nil
}
