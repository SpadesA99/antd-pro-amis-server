/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 15:16:02
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-15 15:33:10
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\getmenurolelogic.go
 */
package logic

import (
	"context"

	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuRoleLogic {
	return &GetMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuRoleLogic) GetMenuRole(in *webapi.GetMenuRoleReq) (*webapi.GetMenuRoleReply, error) {
	MenuRols, err := l.svcCtx.Db.MenuRole.Query().Limit(int(in.PerPage)).Offset((int(in.Page) - 1) * int(in.PerPage)).All(context.Background())
	if err != nil {
		return nil, err
	}
	total, err := l.svcCtx.Db.MenuRole.Query().Count(context.Background())
	if err != nil {
		return nil, err
	}

	list := []*webapi.GetMenuRoleReply_Data_List{}
	if err := copier.Copy(&list, &MenuRols); err != nil {
		return nil, err
	}

	return &webapi.GetMenuRoleReply{Msg: "success", Data: &webapi.GetMenuRoleReply_Data{Total: int32(total), Items: list}}, nil
}
