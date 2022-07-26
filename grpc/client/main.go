package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cectc/hptx-samples/aggregation_svc/api"
)

const (
	aggregationServiceAddress = "aggregation-svc:8000"
)

func main() {
	r := gin.Default()

	conn, err := grpc.Dial(aggregationServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("failed to connect with the given target: %s, err: %v", aggregationServiceAddress, err)
	}

	aggregationSvcClient := api.NewAggregationServiceClient(conn)

	r.GET("/createSoCommit", func(c *gin.Context) {
		_, err := aggregationSvcClient.CreateSoCommit(context.Background(), &empty.Empty{})
		if err == nil {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		} else {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	})

	r.GET("/createSoRollback", func(c *gin.Context) {
		_, err := aggregationSvcClient.CreateSoRollback(context.Background(), &empty.Empty{})
		if err == nil {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		} else {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	})

	r.Run(":8003")
}
