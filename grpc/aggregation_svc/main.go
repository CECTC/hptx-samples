package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cectc/hptx"
	hptxGrpc "github.com/cectc/hptx/pkg/contrib/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cectc/hptx-samples/aggregation_svc/api"
	"github.com/cectc/hptx-samples/aggregation_svc/svc"
	soApi "github.com/cectc/hptx-samples/order_svc/api"
	inventoryApi "github.com/cectc/hptx-samples/product_svc/api"
)

const (
	soServiceAddress        = "order-svc:8002"
	inventoryServiceAddress = "product-svc:8001"
)

func main() {
	configPath := os.Getenv("ConfigPath")
	hptx.InitFromFile(configPath)

	conn1, err := grpc.Dial(soServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect with the given target: %s, err: %v", soServiceAddress, err)
	}
	conn2, err := grpc.Dial(inventoryServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect with the given target: %s, err: %v", inventoryServiceAddress, err)
	}

	soClient := soApi.NewSoServiceClient(conn1)
	inventoryClient := inventoryApi.NewInventoryServiceClient(conn2)

	service := &svc.Service{
		SoClient:        soClient,
		InventoryClient: inventoryClient,
	}

	// Create a gRPC server object
	var s = grpc.NewServer(grpc.ChainUnaryInterceptor(
		hptxGrpc.GlobalTransactionInterceptor([]*hptxGrpc.GlobalTransactionInfo{
			{
				FullMethod: "/api.AggregationService/CreateSoCommit",
				Timeout:    60000,
			},
			{
				FullMethod: "/api.AggregationService/CreateSoRollback",
				Timeout:    60000,
			},
		})))

	// Attach the Greeter service to the server
	api.RegisterAggregationServiceServer(s, service)
	// Serve gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", 8000))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
