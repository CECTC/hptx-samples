package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"gitee.com/chunanyong/zorm"
	"github.com/cectc/hptx"
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/hptx/pkg/resource"
	"github.com/cectc/mysql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cectc/hptx-samples/order_svc/dao"
)

var dbDao *zorm.DBDao

func InitDbByZorm(db *sql.DB) error {
	dbConfig := zorm.DataSourceConfig{
		//连接数据库DSN
		DSN: config.GetATConfig().DSN,
		//数据库驱动名称:mysql,postgres,oci8...
		DriverName: "mysql",
		//数据库类型(方言判断依据):mysql,postgresql,oracle...
		DBType: "mysql",
		//设置慢日志
		SlowSQLMillis: 0,
		//最大连接数 默认50
		MaxOpenConns: 0,
		//最大空闲数 默认50
		MaxIdleConns: 0,
		//连接存活秒时间. 默认600
		ConnMaxLifetimeSecond: 0,
		//事务隔离级别的默认配置,默认为nil
		DefaultTxOptions: nil,
	}
	if db != nil {
		dbConfig.DSN = ""
		dbConfig.SQLDB = db
	}

	var err error
	dbDao, err = zorm.NewDBDao(&dbConfig)
	if err != nil {
		log.Fatalf("数据库连接异常 %v", err)
		return err
	}

	log.Println("数据库连接成功")
	return nil
}

func main() {
	r := gin.Default()

	configPath := os.Getenv("ConfigPath")
	hptx.InitFromFile(configPath)
	mysql.RegisterResource(config.GetATConfig().DSN)
	resource.InitATBranchResource(mysql.GetDataSourceManager())

	sqlDB, err := sql.Open("mysql", config.GetATConfig().DSN)
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(4 * time.Hour)

	err = InitDbByZorm(sqlDB)
	if err != nil {
		panic(err)
	}

	d := &dao.Dao{
		DBDao: dbDao,
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

		xidStr := c.Request.Header.Get("XID")
		log.Printf("xidStr: %s", xidStr)

		_, err := d.CreateSO(
			context.WithValue(
				context.Background(),
				mysql.XID,
				xidStr),
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
