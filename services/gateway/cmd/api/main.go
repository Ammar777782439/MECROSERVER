package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sikozonpc/kitchen/services/gateway/internal/config"
	"github.com/sikozonpc/kitchen/services/gateway/internal/router"
	"github.com/sikozonpc/kitchen/services/gateway/pkg/clients"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize gRPC clients
	grpcClients, err := clients.InitGrpcClients(cfg)
	if err != nil {
		// InitGrpcClients already logs the fatal error, so we just exit.
		return
	}

	app := fiber.New()

	// Simple health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Gateway is running!")
	})

	// Setup API routes
	router.SetupRoutes(app, grpcClients)

	log.Println("Starting API Gateway on port 8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
