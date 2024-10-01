package handlers

import (
	userservice "github.com/gambitier/gorest/services/users"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *userservice.UserService
}

func NewUserHandler(userService *userservice.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// GetAllUsers gets all users.
// @Summary gets all users
// @Description get all users
// @Tags users
// @Accept json
// @Produce json
// @Router /users/ [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	data, err := h.UserService.GetAllUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "couldn't find all users")
	}

	c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Users fetched successfully",
		"data":    data,
	})

	return nil
}
