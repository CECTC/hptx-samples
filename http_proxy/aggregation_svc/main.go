package main

import (
	"os"

	"github.com/cectc/hptx"
	"github.com/cectc/hptx/pkg/tm"
	"github.com/gin-gonic/gin"

	"github.com/cectc/hptx-samples/aggregation_svc/svc"
)

func main() {
	r := gin.Default()

	configPath := os.Getenv("ConfigPath")
	hptx.InitFromFile(configPath)
	tm.Implement(svc.ProxySvc)

	r.GET("/createSoCommit", func(c *gin.Context) {

		svc.ProxySvc.CreateSo(c, false)

		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
		})
	})

	r.GET("/createSoRollback", func(c *gin.Context) {

		svc.ProxySvc.CreateSo(c, true)

		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
		})
	})

	r.Run(":8003")
}
