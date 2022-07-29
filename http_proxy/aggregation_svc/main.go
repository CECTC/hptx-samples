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
		err := svc.ProxySvc.CreateSo(c, false)
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
		err := svc.ProxySvc.CreateSo(c, true)
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
