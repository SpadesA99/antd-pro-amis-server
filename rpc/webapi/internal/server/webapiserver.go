// Code generated by goctl. DO NOT EDIT!
// Source: webapi.proto

package server

import (
	"context"

	"antd-pro-amis-server/rpc/webapi/internal/logic"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"
)

type WebapiServer struct {
	svcCtx *svc.ServiceContext
	webapi.UnimplementedWebapiServer
}

func NewWebapiServer(svcCtx *svc.ServiceContext) *WebapiServer {
	return &WebapiServer{
		svcCtx: svcCtx,
	}
}

// AmisSchema
func (s *WebapiServer) GetAmisSchema(ctx context.Context, in *webapi.GetAmisSchemaReq) (*webapi.GetAmisSchemaReply, error) {
	l := logic.NewGetAmisSchemaLogic(ctx, s.svcCtx)
	return l.GetAmisSchema(in)
}

// 更新AmisSchema或者添加AmisSchema
func (s *WebapiServer) UpdateAmisSchema(ctx context.Context, in *webapi.UpdateAmisSchemaReq) (*webapi.UpdateAmisSchemaReply, error) {
	l := logic.NewUpdateAmisSchemaLogic(ctx, s.svcCtx)
	return l.UpdateAmisSchema(in)
}

// 菜单角色管理
func (s *WebapiServer) AddMenuRole(ctx context.Context, in *webapi.MenuRoleReq) (*webapi.MenuRoleReply, error) {
	l := logic.NewAddMenuRoleLogic(ctx, s.svcCtx)
	return l.AddMenuRole(in)
}

func (s *WebapiServer) DelMenuRole(ctx context.Context, in *webapi.MenuRoleReq) (*webapi.MenuRoleReply, error) {
	l := logic.NewDelMenuRoleLogic(ctx, s.svcCtx)
	return l.DelMenuRole(in)
}

func (s *WebapiServer) GetMenuRole(ctx context.Context, in *webapi.GetMenuRoleReq) (*webapi.GetMenuRoleReply, error) {
	l := logic.NewGetMenuRoleLogic(ctx, s.svcCtx)
	return l.GetMenuRole(in)
}

// 菜单管理
func (s *WebapiServer) AddMenu(ctx context.Context, in *webapi.AddMenuReq) (*webapi.AddMenuReply, error) {
	l := logic.NewAddMenuLogic(ctx, s.svcCtx)
	return l.AddMenu(in)
}

func (s *WebapiServer) GetMenu(ctx context.Context, in *webapi.GetMenuReq) (*webapi.GetMenuReply, error) {
	l := logic.NewGetMenuLogic(ctx, s.svcCtx)
	return l.GetMenu(in)
}

func (s *WebapiServer) DelMenu(ctx context.Context, in *webapi.DelMenuReq) (*webapi.DelMenuReply, error) {
	l := logic.NewDelMenuLogic(ctx, s.svcCtx)
	return l.DelMenu(in)
}

func (s *WebapiServer) UpdateMenu(ctx context.Context, in *webapi.UpdateMenuReq) (*webapi.UpdateMenuReply, error) {
	l := logic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

func (s *WebapiServer) BulkDelMenu(ctx context.Context, in *webapi.BulkDelMenuReq) (*webapi.DelMenuReply, error) {
	l := logic.NewBulkDelMenuLogic(ctx, s.svcCtx)
	return l.BulkDelMenu(in)
}

func (s *WebapiServer) OrderMenu(ctx context.Context, in *webapi.OrderMenuReq) (*webapi.OrderMenuReply, error) {
	l := logic.NewOrderMenuLogic(ctx, s.svcCtx)
	return l.OrderMenu(in)
}

func (s *WebapiServer) GetMenuTree(ctx context.Context, in *webapi.GetMenuTreeReq) (*webapi.GetMenuTreeReply, error) {
	l := logic.NewGetMenuTreeLogic(ctx, s.svcCtx)
	return l.GetMenuTree(in)
}