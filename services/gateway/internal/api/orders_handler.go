package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikozonpc/kitchen/services/common/genproto/orders"
)

// OrdersHandler handles API requests related to orders.
type OrdersHandler struct {
	ordersClient orders.OrderServiceClient
}

// NewOrdersHandler creates a new OrdersHandler.
func NewOrdersHandler(ordersClient orders.OrderServiceClient) *OrdersHandler {
	return &OrdersHandler{ordersClient: ordersClient}
}

// CreateOrder is the handler for creating a new order.
func (h *OrdersHandler) CreateOrder(c *fiber.Ctx) error {
	// Define a struct to parse the request body
	type createOrderPayload struct {
		CustomerID int32 `json:"customerID"`
		ProductID  int32 `json:"productID"`
		Quantity   int32 `json:"quantity"`
	}

	payload := new(createOrderPayload)

	// Parse the request body into the payload struct
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Basic validation
	if payload.CustomerID == 0 || payload.ProductID == 0 || payload.Quantity <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Missing required fields or invalid quantity")
	}

	// Call the gRPC service
	resp, err := h.ordersClient.CreateOrder(c.Context(), &orders.CreateOrderRequest{
		CustomerID: payload.CustomerID,
		ProductID:  payload.ProductID,
		Quantity:   payload.Quantity,
	})

	if err != nil {
		// In a real app, you'd inspect the gRPC error and return a more specific HTTP status
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create order")
	}

	return c.JSON(resp)
}
