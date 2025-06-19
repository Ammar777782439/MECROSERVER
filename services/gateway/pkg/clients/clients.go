package clients

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sikozonpc/kitchen/services/common/genproto/orders"
	// We will add the kitchen proto import here later
	"github.com/sikozonpc/kitchen/services/gateway/internal/config"
)

// GrpcClients holds the gRPC clients for all microservices.
type GrpcClients struct {
	Orders orders.OrderServiceClient
	// Kitchen kitchen.KitchenServiceClient
}

// InitGrpcClients initializes and returns gRPC clients for the services.
func InitGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	// Establish connection to the orders service
	ordersConn, err := grpc.Dial(cfg.OrdersSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to orders service: %v", err)
		return nil, err
	}

	// Establish connection to the kitchen service
	// kitchenConn, err := grpc.Dial(cfg.KitchenSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Failed to connect to kitchen service: %v", err)
	// 	return nil, err
	// }

	log.Println("Successfully connected to gRPC services")

	return &GrpcClients{
		Orders: orders.NewOrderServiceClient(ordersConn),
		// Kitchen: kitchen.NewKitchenServiceClient(kitchenConn),
	}, nil
}
