/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-20 14:50:52
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-20 14:53:23
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\bulkdelmenulogic.go
 */
package logic

import (
	"context"

	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/utils"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type BulkDelMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBulkDelMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BulkDelMenuLogic {
	return &BulkDelMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BulkDelMenuLogic) BulkDelMenu(in *webapi.BulkDelMenuReq) (*webapi.DelMenuReply, error) {
	ids := utils.ToIntArray(in.Ids)

	if _, err := l.svcCtx.Db.Menu.Delete().Where(menu.IDIn(ids...)).Exec(context.Background()); err != nil {
		return nil, err
	}

	return &webapi.DelMenuReply{Msg: "success"}, nil
}
