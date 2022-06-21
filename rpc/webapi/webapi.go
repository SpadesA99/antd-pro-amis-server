/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 11:17:15
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-15 11:44:17
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\webapi.go
 */
package main

import (
	"flag"
	"fmt"

	"antd-pro-amis-server/rpc/webapi/internal/config"
	"antd-pro-amis-server/rpc/webapi/internal/server"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/webapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewWebapiServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		webapi.RegisterWebapiServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
