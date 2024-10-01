//go:generate swag init -o ./_apidocs

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	appcontext "github.com/gambitier/gorest/app-context"
	httpserver "github.com/gambitier/gorest/http-server"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	appCtx, err := appcontext.InitAppContext()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	server := httpserver.NewHttpServer(appCtx)
	server.Configure()
	server.RegisterRoutes()

	// Channel to listen for errors
	errChan := make(chan error, 1)

	// Start Fiber server in a goroutine
	go func() {
		if err := server.RunServer(3000); err != nil {
			errChan <- fmt.Errorf("failed to start Fiber server: %w", err)
		}
	}()

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("server error: %v", err)
	case <-quit:
		log.Println("shutting down servers...")

		// Gracefully shut down the Fiber server
		if err := server.Shutdown(); err != nil {
			log.Fatalf("Fiber server shutdown error: %v", err)
		}

		log.Println("servers shut down gracefully")
	}

}
