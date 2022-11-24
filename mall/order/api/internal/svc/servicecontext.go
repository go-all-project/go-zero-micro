package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero/mall/order/api/internal/config"
	"go_zero/mall/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	//配置文件上下文及需要使用的 RPC 上下文
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
