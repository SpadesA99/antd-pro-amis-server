/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 15:16:02
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-16 18:29:17
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\delmenurolelogic.go
 */
package logic

import (
	"context"

	"antd-pro-amis-server/rpc/ent/menurole"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMenuRoleLogic {
	return &DelMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMenuRoleLogic) DelMenuRole(in *webapi.MenuRoleReq) (*webapi.MenuRoleReply, error) {
	if _, err := l.svcCtx.Db.MenuRole.Delete().Where(menurole.RoleName(in.RoleName)).Exec(context.Background()); err != nil {
		return nil, err
	}

	return &webapi.MenuRoleReply{Msg: "success"}, nil
}
