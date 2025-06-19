package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sikozonpc/kitchen/services/gateway/internal/api"
	"github.com/sikozonpc/kitchen/services/gateway/pkg/clients"
)

// SetupRoutes configures the API routes for the gateway.
func SetupRoutes(app *fiber.App, grpcClients *clients.GrpcClients) {
	// Add a logger middleware to log all incoming requests
	app.Use(logger.New())

	// Create handlers
	ordersHandler := api.NewOrdersHandler(grpcClients.Orders)

	// Group routes under /api/v1
	apiV1 := app.Group("/api/v1")

	// Orders routes
	ordersRoutes := apiV1.Group("/orders")
	ordersRoutes.Post("/", ordersHandler.CreateOrder)

	// You can add other routes for kitchen, etc. here
}
