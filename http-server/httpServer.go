package httpserver

import (
	"fmt"
	"log"

	"github.com/gambitier/gorest/api/users/routes"
	appcontext "github.com/gambitier/gorest/app-context"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HttpServer struct {
	app        *fiber.App
	appContext *appcontext.AppContext
}

func NewHttpServer(appContext *appcontext.AppContext) *HttpServer {
	return &HttpServer{
		app:        fiber.New(),
		appContext: appContext,
	}
}

func (server *HttpServer) Configure() {
	server.app.Use(etag.New())
	server.app.Use(cache.New())
	server.app.Use(compress.New())
	server.app.Use(recover.New())
	server.app.Use(idempotency.New())
	server.app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/swagger" {
			return c.Next()
		}
		return helmet.New()(c)
	})

	// Middleware
	server.app.Use(logger.New())

	// Config swagger
	server.app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./_apidocs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}))
}

func (server *HttpServer) RunServer(port int) error {
	address := fmt.Sprintf(":%v", port)

	log.Printf("Starting server on %s\n", address)

	if err := server.app.Listen(address); err != nil {
		return err
	}

	return nil
}

func (server *HttpServer) Shutdown() error {
	return server.app.Shutdown()
}

func (server *HttpServer) RegisterRoutes() {
	routes.RegisterUserRoutes(server.appContext, server.app)
}
