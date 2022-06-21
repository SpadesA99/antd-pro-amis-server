/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-16 14:23:30
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-21 13:43:30
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\addmenulogic.go
 */
package logic

import (
	"context"
	"strings"
	"time"

	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/utils"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 菜单管理
func (l *AddMenuLogic) AddMenu(in *webapi.AddMenuReq) (*webapi.AddMenuReply, error) {
	if m, err := l.svcCtx.Db.Menu.Query().Where(menu.Path(in.Path)).Exist(context.Background()); err != nil {
		return nil, err
	} else {
		if m {
			return &webapi.AddMenuReply{Status: 400, Msg: "路由已存在"}, nil
		}
	}

	if _, err := l.svcCtx.Db.Menu.Create().
		SetMenuName(in.Name).
		SetPath(in.Path).
		SetIcon(in.Icon).
		SetParentID(in.ParentId).
		SetRoles(strings.Split(in.Roles, ",")).
		SetLayout(utils.IfThen(in.Layout == 1, true, false)).
		SetHideInMenu(utils.IfThen(in.HideInMenu == 1, true, false)).
		SetSort(time.Now().Unix()).
		Save(context.Background()); err != nil {
		return nil, err
	}

	return &webapi.AddMenuReply{Status: 0, Msg: "创建菜单成功"}, nil
}
