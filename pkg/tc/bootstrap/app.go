package bootstrap

import (
	"log"
	"strconv"

	gxnet "github.com/dubbogo/gost/net"
	"github.com/duolacloud/seata-golang/pkg/base/common"
	tx_config "github.com/duolacloud/seata-golang/pkg/tc/config"
	"github.com/duolacloud/seata-golang/pkg/tc/holder"
	"github.com/duolacloud/seata-golang/pkg/tc/lock"
	_ "github.com/duolacloud/seata-golang/pkg/tc/metrics"
	"github.com/duolacloud/seata-golang/pkg/tc/server"
	"github.com/duolacloud/seata-golang/pkg/util/uuid"
	"github.com/micro/go-micro/v2/config"
)

func StartServer(config config.Config) error {
	ip, _ := gxnet.GetLocalIP()

	serverNode := config.Get("tc", "server_node").Int(1)

	var serverConfig tx_config.ServerConfig
	config.Get("tc").Scan(&serverConfig)

	tx_config.SetServerConfig(serverConfig)
	port, _ := strconv.Atoi(serverConfig.Port)
	common.XID.Init(ip, port)

	log.Printf("start tc-server, addr: %s:%d, node: %d", serverConfig.Host, port, serverNode)

	uuid.Init(serverNode)
	lock.Init()
	holder.Init()
	srv := server.NewServer()
	srv.Start(serverConfig.Host + ":" + serverConfig.Port)
	return nil
}
