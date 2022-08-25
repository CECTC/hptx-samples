package main

import (
	"os"

	"github.com/cectc/hptx"
	hptxGin "github.com/cectc/hptx/pkg/contrib/gin"
	"github.com/gin-gonic/gin"

	"github.com/cectc/hptx-samples/aggregation_svc/svc"
)

func main() {
	r := gin.Default()

	configPath := os.Getenv("ConfigPath")
	hptx.InitFromFile(configPath)

	r.GET("/createSoCommit", hptxGin.GlobalTransaction(60000), func(c *gin.Context) {
		err := svc.Service.CreateSo(c, false)
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

	r.GET("/createSoRollback", hptxGin.GlobalTransaction(60000), func(c *gin.Context) {
		err := svc.Service.CreateSo(c, true)
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
