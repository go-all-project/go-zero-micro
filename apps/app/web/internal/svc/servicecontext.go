package svc

import (
	"context"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero/apps/app/web/internal/config"
	"go_zero/apps/app/web/internal/middleware"
	"go_zero/apps/order/rpc/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config config.Config
	// 在这里注册BFF中的 RPC 服务
	OrderRPCClient order.Order

	// 中间件
	login rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		OrderRPCClient: order.NewOrder(zrpc.MustNewClient(c.OrderRPC, zrpc.WithUnaryClientInterceptor(testClientInterceptor))),
		login:          middleware.NewLoginMiddleware().Handle,
	}
}

// 客户端 rpc 拦截器
func testClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	// 通过 metadata 传参
	md := metadata.New(map[string]string{"username": "zhang-san"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	if err := invoker(ctx, method, req, reply, cc, opts...); err != nil {
		return err
	}

	return nil
}
