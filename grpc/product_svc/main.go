package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/cectc/hptx"
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/mysql"
	"google.golang.org/grpc"

	"github.com/cectc/hptx-samples/product_svc/api"
	"github.com/cectc/hptx-samples/product_svc/dao"
	"github.com/cectc/hptx-samples/product_svc/svc"
)

func main() {
	configPath := os.Getenv("ConfigPath")
	hptx.InitFromFile(configPath)
	mysql.RegisterATResource(config.GetATConfig().DSN)

	sqlDB, err := sql.Open("mysql", config.GetATConfig().DSN)
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(4 * time.Hour)

	d := &dao.Dao{
		DB: sqlDB,
	}

	service := &svc.Service{Dao: d}

	// Create a gRPC server object
	var s = grpc.NewServer()
	// Attach the Greeter service to the server
	api.RegisterInventoryServiceServer(s, service)
	// Serve gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", 8001))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
