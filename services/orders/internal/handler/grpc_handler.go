package handler

import (
	"context"

	pb "github.com/sikozonpc/kitchen/services/common/genproto/orders"
	"github.com/sikozonpc/kitchen/services/orders/internal/service"
)

// GrpcHandler is the gRPC handler for the orders service.
type GrpcHandler struct {
	pb.UnimplementedOrderServiceServer // Embed for forward compatibility
	service *service.OrderService
}

// NewGrpcHandler creates a new GrpcHandler.
func NewGrpcHandler(service *service.OrderService) *GrpcHandler {
	return &GrpcHandler{service: service}
}

// CreateOrder handles the gRPC request to create an order.
func (h *GrpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	input := service.CreateOrderInput{
		CustomerID: uint(req.CustomerID),
		ProductID:  uint(req.ProductID),
		Quantity:   int(req.Quantity),
	}

		_, err := h.service.CreateOrder(ctx, input)
	if err != nil {
		// In a real app, you would map domain errors to gRPC status codes
		return nil, err
	}

		return &pb.CreateOrderResponse{
		Status: "CREATED",
	}, nil
}
