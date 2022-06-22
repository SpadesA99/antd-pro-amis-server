// Code generated by goctl. DO NOT EDIT!
// Source: webapi.proto

package webapi

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Webapi interface {
		// AmisSchema
		GetAmisSchema(ctx context.Context, in *GetAmisSchemaReq, opts ...grpc.CallOption) (*GetAmisSchemaReply, error)
		// 更新AmisSchema或者添加AmisSchema
		UpdateAmisSchema(ctx context.Context, in *UpdateAmisSchemaReq, opts ...grpc.CallOption) (*UpdateAmisSchemaReply, error)
		// 菜单角色管理
		AddMenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleReply, error)
		DelMenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleReply, error)
		GetMenuRole(ctx context.Context, in *GetMenuRoleReq, opts ...grpc.CallOption) (*GetMenuRoleReply, error)
		// 菜单管理
		AddMenu(ctx context.Context, in *AddMenuReq, opts ...grpc.CallOption) (*AddMenuReply, error)
		GetMenu(ctx context.Context, in *GetMenuReq, opts ...grpc.CallOption) (*GetMenuReply, error)
		DelMenu(ctx context.Context, in *DelMenuReq, opts ...grpc.CallOption) (*DelMenuReply, error)
		UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuReply, error)
		SortMenu(ctx context.Context, in *SortMenuReq, opts ...grpc.CallOption) (*SortMenuReply, error)
		BulkDelMenu(ctx context.Context, in *BulkDelMenuReq, opts ...grpc.CallOption) (*DelMenuReply, error)
		OrderMenu(ctx context.Context, in *OrderMenuReq, opts ...grpc.CallOption) (*OrderMenuReply, error)
		GetMenuTree(ctx context.Context, in *GetMenuTreeReq, opts ...grpc.CallOption) (*GetMenuTreeReply, error)
	}

	defaultWebapi struct {
		cli zrpc.Client
	}
)

func NewWebapi(cli zrpc.Client) Webapi {
	return &defaultWebapi{
		cli: cli,
	}
}

// AmisSchema
func (m *defaultWebapi) GetAmisSchema(ctx context.Context, in *GetAmisSchemaReq, opts ...grpc.CallOption) (*GetAmisSchemaReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.GetAmisSchema(ctx, in, opts...)
}

// 更新AmisSchema或者添加AmisSchema
func (m *defaultWebapi) UpdateAmisSchema(ctx context.Context, in *UpdateAmisSchemaReq, opts ...grpc.CallOption) (*UpdateAmisSchemaReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.UpdateAmisSchema(ctx, in, opts...)
}

// 菜单角色管理
func (m *defaultWebapi) AddMenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.AddMenuRole(ctx, in, opts...)
}

func (m *defaultWebapi) DelMenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.DelMenuRole(ctx, in, opts...)
}

func (m *defaultWebapi) GetMenuRole(ctx context.Context, in *GetMenuRoleReq, opts ...grpc.CallOption) (*GetMenuRoleReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.GetMenuRole(ctx, in, opts...)
}

// 菜单管理
func (m *defaultWebapi) AddMenu(ctx context.Context, in *AddMenuReq, opts ...grpc.CallOption) (*AddMenuReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.AddMenu(ctx, in, opts...)
}

func (m *defaultWebapi) GetMenu(ctx context.Context, in *GetMenuReq, opts ...grpc.CallOption) (*GetMenuReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.GetMenu(ctx, in, opts...)
}

func (m *defaultWebapi) DelMenu(ctx context.Context, in *DelMenuReq, opts ...grpc.CallOption) (*DelMenuReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.DelMenu(ctx, in, opts...)
}

func (m *defaultWebapi) UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...grpc.CallOption) (*UpdateMenuReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

func (m *defaultWebapi) SortMenu(ctx context.Context, in *SortMenuReq, opts ...grpc.CallOption) (*SortMenuReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.SortMenu(ctx, in, opts...)
}

func (m *defaultWebapi) BulkDelMenu(ctx context.Context, in *BulkDelMenuReq, opts ...grpc.CallOption) (*DelMenuReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.BulkDelMenu(ctx, in, opts...)
}

func (m *defaultWebapi) OrderMenu(ctx context.Context, in *OrderMenuReq, opts ...grpc.CallOption) (*OrderMenuReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.OrderMenu(ctx, in, opts...)
}

func (m *defaultWebapi) GetMenuTree(ctx context.Context, in *GetMenuTreeReq, opts ...grpc.CallOption) (*GetMenuTreeReply, error) {
	client := NewWebapiClient(m.cli.Conn())
	return client.GetMenuTree(ctx, in, opts...)
}
