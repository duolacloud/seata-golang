package rpc_client

import "github.com/duolacloud/seata-golang/pkg/base/protocal"

type RpcRMMessage struct {
	RpcMessage    protocal.RpcMessage
	ServerAddress string
}
