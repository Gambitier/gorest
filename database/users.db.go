package database

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
)

// UserData represents a user record in our "database"
type UserData struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserRepoHandler contains the user table and the mutex for thread-safe operations
type UserRepoHandler struct {
	users []UserData
	mu    sync.Mutex
}

// GenerateId generates a random ID for new users (as a string)
func GenerateId() string {
	return strconv.Itoa(rand.Intn(10000)) // Generates a random number as a string
}

// Initialize the database with some values
func NewUserRepoHandler() *UserRepoHandler {
	handler := &UserRepoHandler{}
	handler.mu.Lock()
	defer handler.mu.Unlock()

	handler.users = []UserData{
		{Id: GenerateId(), FirstName: "John", LastName: "Doe"},
		{Id: GenerateId(), FirstName: "Jane", LastName: "Smith"},
	}

	return handler
}

// AddUser adds a new user to the "database" in a thread-safe manner
func (h *UserRepoHandler) AddUser(firstName string, lastName string) UserData {
	h.mu.Lock() // Lock the table for writing
	defer h.mu.Unlock()

	user := UserData{
		Id:        GenerateId(),
		FirstName: firstName,
		LastName:  lastName,
	}
	h.users = append(h.users, user)
	return user
}

// GetUserById retrieves a user by their ID in a thread-safe manner
func (h *UserRepoHandler) GetUserById(id string) (UserData, error) {
	h.mu.Lock() // Lock the table for reading
	defer h.mu.Unlock()

	for _, user := range h.users {
		if user.Id == id {
			return user, nil
		}
	}
	return UserData{}, errors.New("user not found")
}

// GetAllUsers returns all users in the "database" in a thread-safe manner
func (h *UserRepoHandler) GetAllUsers() []UserData {
	h.mu.Lock() // Lock the table for reading
	defer h.mu.Unlock()

	return h.users
}

// UpdateUser updates an existing user in the "database" in a thread-safe manner
func (h *UserRepoHandler) UpdateUser(id string, firstName string, lastName string) (UserData, error) {
	h.mu.Lock() // Lock the table for writing
	defer h.mu.Unlock()

	for i, user := range h.users {
		if user.Id == id {
			h.users[i].FirstName = firstName
			h.users[i].LastName = lastName
			return h.users[i], nil
		}
	}
	return UserData{}, errors.New("user not found")
}

// DeleteUser deletes a user from the "database" in a thread-safe manner
func (h *UserRepoHandler) DeleteUser(id string) error {
	h.mu.Lock() // Lock the table for writing
	defer h.mu.Unlock()

	for i, user := range h.users {
		if user.Id == id {
			h.users = append(h.users[:i], h.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
