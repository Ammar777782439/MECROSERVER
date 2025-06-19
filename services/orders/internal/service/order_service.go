package service

import (
	"context"

	"github.com/sikozonpc/kitchen/services/orders/internal/domain"
	"github.com/sikozonpc/kitchen/services/orders/internal/store"
)

// OrderService provides business logic for orders.
type OrderService struct {
	store store.OrderStore
}

// NewOrderService creates a new OrderService.
func NewOrderService(store store.OrderStore) *OrderService {
	return &OrderService{store: store}
}

// CreateOrderInput defines the input for creating an order.
type CreateOrderInput struct {
	CustomerID uint `json:"customerID"`
	ProductID  uint `json:"productID"`
	Quantity   int  `json:"quantity"`
}

// CreateOrder handles the business logic for creating a new order.
func (s *OrderService) CreateOrder(ctx context.Context, input CreateOrderInput) (*domain.Order, error) {
	// Here you could add more complex business logic, e.g., checking customer existence,
	// product availability, etc.

	order := &domain.Order{
		CustomerID: input.CustomerID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
	}

	err := s.store.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
