package main

import (
	"database/sql"
	"log"
	"os"

	"gitee.com/chunanyong/zorm"
	"github.com/cectc/hptx"
	"github.com/cectc/hptx/pkg/tm"
	"github.com/gin-gonic/gin"

	"github.com/cectc/hptx-samples/aggregation_svc/svc"
	_ "github.com/go-sql-driver/mysql"
)

var dbDao *zorm.DBDao

func InitDbByZorm(db *sql.DB) error {
	dbConfig := zorm.DataSourceConfig{
		//连接数据库DSN
		DSN: "root:123456@tcp(hptx-mysql:3306)/order?timeout=10s&readTimeout=10s&writeTimeout=10s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		//数据库类型
		DriverName: "mysql",
		DBType:     "mysql",
		//是否的打印sql
		// PrintSQL: true,
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
	tm.Implement(svc.ProxySvc)

	err := InitDbByZorm(nil)
	if err != nil {
		panic(err)
	}

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
