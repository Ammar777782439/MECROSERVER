package config

import (
	"log"
	"os"
)

// Config holds the application's configuration.
type Config struct {
	OrdersSvcUrl  string
	KitchenSvcUrl string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() *Config {
	ordersSvcUrl := os.Getenv("ORDERS_SERVICE_URL")
	if ordersSvcUrl == "" {
		// Default value for local development
		ordersSvcUrl = "localhost:9091"
		log.Printf("ORDERS_SERVICE_URL not set, using default: %s", ordersSvcUrl)
	}

	kitchenSvcUrl := os.Getenv("KITCHEN_SERVICE_URL")
	if kitchenSvcUrl == "" {
		// Default value for local development
		kitchenSvcUrl = "localhost:50052"
		log.Printf("KITCHEN_SERVICE_URL not set, using default: %s", kitchenSvcUrl)
	}

	return &Config{
		OrdersSvcUrl:  ordersSvcUrl,
		KitchenSvcUrl: kitchenSvcUrl,
	}
}
