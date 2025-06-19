package store

import (
	"context"

	"github.com/sikozonpc/kitchen/services/orders/internal/domain"
)

// OrderStore defines the interface for interacting with order data.
type OrderStore interface {
	Create(ctx context.Context, order *domain.Order) error
}
