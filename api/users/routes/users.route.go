package routes

import (
	"github.com/gambitier/gorest/api/users/handlers"
	appcontext "github.com/gambitier/gorest/app-context"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *appcontext.AppContext, httpServer *fiber.App) {
	userHandler := handlers.NewUserHandler(app.Services.UserService)
	users := httpServer.Group("/users")
	// users.Post("/")
	// users.Post("/avatar", userHandler.UploadAvatar)
	users.Get("/", userHandler.GetAllUsers)
}
