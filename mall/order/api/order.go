package main

import (
	"flag"
	"fmt"

	"go_zero/mall/order/api/internal/config"
	"go_zero/mall/order/api/internal/handler"
	"go_zero/mall/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	// 加载配置文件
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 根据配置创建服务器
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 创建服务的上下文对象
	ctx := svc.NewServiceContext(c)

	// 注册路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
