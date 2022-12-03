package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero/apps/app/web/internal/config"
	"go_zero/apps/app/web/internal/middleware"
	"go_zero/apps/order/rpc/order"
)

type ServiceContext struct {
	Config config.Config
	// 在这里注册BFF中的 RPC 服务
	OrderRPC order.Order

	// 中间件
	login rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRPC: order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		login:    middleware.NewLoginMiddleware().Handle,
	}
}
