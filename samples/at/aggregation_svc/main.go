package main

import (
	"github.com/gin-gonic/gin"
)

import (
	"github.com/duolacloud/seata-golang/pkg/client"
	"github.com/duolacloud/seata-golang/pkg/client/config"
	"github.com/duolacloud/seata-golang/pkg/client/tm"
	"github.com/duolacloud/seata-golang/samples/at/aggregation_svc/svc"
)

var configPath = "/Users/scottlewis/dksl/git/1/seata-golang/samples/at/aggregation_svc/conf/client.yml"

func main() {
	r := gin.Default()
	config.InitConf(configPath)
	client.NewRpcClient()
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
