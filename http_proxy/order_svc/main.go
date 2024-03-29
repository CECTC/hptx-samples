package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/cectc/hptx"
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/hptx/pkg/constant"
	"github.com/cectc/mysql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cectc/hptx-samples/order_svc/dao"
)

func main() {
	r := gin.Default()

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

	r.POST("/createSo", func(c *gin.Context) {
		type req struct {
			Req []*dao.SoMaster
		}
		var q req
		if err := c.ShouldBindJSON(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := d.CreateSO(
			context.WithValue(
				context.Background(),
				constant.XID,
				c.Request.Header.Get("XID")),
			q.Req)

		if err != nil {
			logrus.Error(err)
			c.JSON(400, gin.H{
				"success": false,
				"message": "fail",
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		}
	})
	r.Run(":8002")
}
