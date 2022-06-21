/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-16 18:29:09
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-20 14:06:45
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\delmenulogic.go
 */
package logic

import (
	"context"

	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMenuLogic {
	return &DelMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMenuLogic) DelMenu(in *webapi.DelMenuReq) (*webapi.DelMenuReply, error) {
	if _, err := l.svcCtx.Db.Menu.Delete().Where(menu.ID(int(in.Id))).Exec(context.Background()); err != nil {
		return nil, err
	}

	return &webapi.DelMenuReply{Msg: "success"}, nil
}
