package config

import "github.com/zeromicro/go-zero/zrpc"

type MySQL struct {
	DNS string
}

type Config struct {
	zrpc.RpcServerConf
	MySQL
}
