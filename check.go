package main

import (
	"flag"
	"fmt"

	"github.com/lokichoggio/check/check"
	"github.com/lokichoggio/check/internal/config"
	"github.com/lokichoggio/check/internal/server"
	"github.com/lokichoggio/check/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/dev/add.yaml", "the config file")
var listenOn = flag.String("listen", "", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	if *listenOn != "" {
		c.ListenOn = *listenOn
	}

	ctx := svc.NewServiceContext(c)
	srv := server.NewCheckerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		check.RegisterCheckerServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
