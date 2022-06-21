/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 15:16:02
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-15 15:17:53
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\addmenurolelogic.go
 */
package logic

import (
	"context"
	"errors"

	"antd-pro-amis-server/rpc/ent/menurole"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuRoleLogic {
	return &AddMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 菜单角色管理
func (l *AddMenuRoleLogic) AddMenuRole(in *webapi.MenuRoleReq) (*webapi.MenuRoleReply, error) {
	if l.svcCtx.Db.MenuRole.Query().Where(menurole.RoleName(in.RoleName)).ExistX(context.Background()) {
		return nil, errors.New("角色名称已存在")
	}

	if _, err := l.svcCtx.Db.MenuRole.Create().SetRoleName(in.RoleName).Save(context.Background()); err != nil {
		return nil, err
	}

	return &webapi.MenuRoleReply{Msg: "success"}, nil
}
