package store

import (
	"context"

	"github.com/sikozonpc/kitchen/services/orders/internal/domain"
	"gorm.io/gorm"
)

// GormOrderStore is an implementation of OrderStore using GORM.
type GormOrderStore struct {
	db *gorm.DB
}

// NewGormOrderStore creates a new GormOrderStore.
func NewGormOrderStore(db *gorm.DB) *GormOrderStore {
	return &GormOrderStore{db: db}
}

// Create creates a new order in the database.
func (s *GormOrderStore) Create(ctx context.Context, order *domain.Order) error {
	return s.db.WithContext(ctx).Create(order).Error
}
