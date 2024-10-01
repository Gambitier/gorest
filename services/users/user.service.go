package userservice

import (
	"fmt"

	"github.com/gambitier/gorest/database"
)

// UserService holds the reference to the database (in-memory user table)
type UserService struct {
	userRepo *database.UserRepoHandler // Inject the database here
}

// NewUserService creates a new UserService and injects the database
func NewUserService(userRepo *database.UserRepoHandler) *UserService {
	return &UserService{
		userRepo,
	}
}

// GetMessage is a simple service method that returns a greeting message
func (s *UserService) GetMessage(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}

// AddUserService adds a user to the database via the UserService
func (s *UserService) AddUser(firstName, lastName string) {
	user := s.userRepo.AddUser(firstName, lastName) // Call the database's AddUser method
	fmt.Printf("User added: %+v\n", user)
}

// GetUserService gets a user by their ID via the UserService
func (s *UserService) GetUser(id string) {
	user, err := s.userRepo.GetUserById(id) // Call the database's GetUserById method
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User found: %+v\n", user)
	}
}

// DeleteUserService deletes a user by their ID via the UserService
func (s *UserService) DeleteUser(id string) {
	err := s.userRepo.DeleteUser(id) // Call the database's DeleteUser method
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User deleted")
	}
}

// GetAllUsersService returns all users via the UserService
func (s *UserService) GetAllUsers() ([]database.UserData, error) {
	users := s.userRepo.GetAllUsers() // Call the database's GetAllUsers method
	return users, nil
}
