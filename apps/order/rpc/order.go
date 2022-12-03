package main

import (
	"context"
	"flag"
	"fmt"

	"go_zero/apps/order/rpc/internal/config"
	"go_zero/apps/order/rpc/internal/server"
	"go_zero/apps/order/rpc/internal/svc"
	"go_zero/apps/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(testInterceptors)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func testInterceptors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	fmt.Printf("testInterceptors info %v", info)
	fmt.Printf("testInterceptors req %v", req)

	resp, err = handler(ctx, req)

	return resp, err
}
