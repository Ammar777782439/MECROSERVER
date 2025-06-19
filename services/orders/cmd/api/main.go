package main

import (
	"log"
	"net"

	pb "github.com/sikozonpc/kitchen/services/common/genproto/orders"
	"github.com/sikozonpc/kitchen/services/orders/internal/domain"
	"github.com/sikozonpc/kitchen/services/orders/internal/handler"
	"github.com/sikozonpc/kitchen/services/orders/internal/service"
	"github.com/sikozonpc/kitchen/services/orders/internal/store"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// TODO: Use a proper config management solution
	dsn := "host=localhost user=postgres password=311 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&domain.Order{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Dependency Injection
	orderStore := store.NewGormOrderStore(db)
	orderService := service.NewOrderService(orderStore)
	grpcHandler := handler.NewGrpcHandler(orderService)

	// Start gRPC Server
	grpcAddr := ":9091" // Same port as before
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, grpcHandler)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
