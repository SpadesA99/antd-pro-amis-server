/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 11:17:15
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-15 11:51:28
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\svc\servicecontext.go
 */
package svc

import (
	"antd-pro-amis-server/rpc/ent"
	"antd-pro-amis-server/rpc/webapi/internal/config"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	Db     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	m, err := ent.Open("mysql", c.DataSource)
	if err != nil {
		logx.Errorf("failed open mysql connect :%v", err)
	}
	if err := m.Schema.Create(context.Background()); err != nil {
		logx.Errorf("failed creating schema resources: %v", err)
	}
	return &ServiceContext{
		Config: c,
		Db:     m,
	}
}
