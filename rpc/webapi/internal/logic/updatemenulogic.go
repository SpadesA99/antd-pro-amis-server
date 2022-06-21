/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-16 18:29:09
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-21 13:39:15
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\updatemenulogic.go
 */
package logic

import (
	"context"
	"errors"
	"strings"

	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/utils"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *webapi.UpdateMenuReq) (*webapi.UpdateMenuReply, error) {
	if b := l.svcCtx.Db.Menu.Query().Where(menu.ID(int(in.Id))).ExistX(context.Background()); !b {
		return nil, errors.New("menu not exist")
	}

	if err := l.svcCtx.Db.Menu.Update().
		Where(menu.ID(int(in.Id))).
		SetMenuName(in.MenuName).
		SetPath(in.Path).
		SetIcon(in.Icon).
		SetParentID(in.ParentId).
		SetStatus(utils.IfThen(in.Status == 1, true, false)).
		SetRoles(strings.Split(in.Roles, ",")).
		SetLayout(utils.IfThen(in.Layout == 1, true, false)).
		SetHideInMenu(utils.IfThen(in.HideInMenu == 1, true, false)).
		Exec(context.Background()); err != nil {
		return nil, err
	}

	return &webapi.UpdateMenuReply{Msg: "success"}, nil
}
