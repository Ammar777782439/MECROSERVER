package domain

import "time"

// Order represents the model for an order
type Order struct {
	ID        uint      `gorm:"primaryKey"`
	CustomerID uint      `json:"customerID"`
	ProductID  uint      `json:"productID"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
