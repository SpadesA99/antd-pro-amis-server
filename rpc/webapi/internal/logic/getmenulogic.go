/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-16 15:51:48
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-21 13:42:22
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\getmenulogic.go
 */
package logic

import (
	"context"
	"strings"

	"antd-pro-amis-server/rpc/ent"
	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/utils"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuLogic) GetMenu(in *webapi.GetMenuReq) (*webapi.GetMenuReply, error) {
	db := l.svcCtx.Db.Menu.Query()

	if in.Id != -1 {
		db = db.Where(menu.ID(int(in.Id)))
	}

	if in.ParentId != -1 {
		db = db.Where(menu.ParentID(in.ParentId))
	}

	Menu, err := db.Order(ent.Asc(menu.FieldSort)).Limit(int(in.PerPage)).Offset((int(in.Page) - 1) * int(in.PerPage)).All(context.Background())
	if err != nil {
		return nil, err
	}
	total, err := l.svcCtx.Db.Menu.Query().Count(context.Background())
	if err != nil {
		return nil, err
	}

	list := make([]*webapi.GetMenuReply_Data_List, len(Menu))
	for i := 0; i < len(list); i++ {
		list[i] = &webapi.GetMenuReply_Data_List{}
		list[i].Roles = strings.Join(Menu[i].Roles, ",")
		list[i].Id = int32(Menu[i].ID)
		list[i].MenuName = Menu[i].MenuName
		list[i].Path = Menu[i].Path
		list[i].ParentId = Menu[i].ParentID
		list[i].Icon = Menu[i].Icon
		list[i].Status = int32(utils.IfThen(Menu[i].Status, 1, 0))
		list[i].Layout = int32(utils.IfThen(Menu[i].Layout, 1, 0))
		list[i].HideInMenu = int32(utils.IfThen(Menu[i].HideInMenu, 1, 0))
	}
	return &webapi.GetMenuReply{Msg: "success", Data: &webapi.GetMenuReply_Data{Total: int32(total), Items: list}}, nil
}
