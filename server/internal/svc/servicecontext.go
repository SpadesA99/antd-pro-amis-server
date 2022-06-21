/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 10:46:37
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-15 11:09:43
 * @FilePath     : \antd-pro-amis-server\server\internal\svc\servicecontext.go
 */
package svc

import (
	"antd-pro-amis-server/rpc/webapi/webapi"
	"antd-pro-amis-server/server/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Webapi webapi.Webapi
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Webapi: webapi.NewWebapi(zrpc.MustNewClient(c.WebapiDsn)),
	}
}
