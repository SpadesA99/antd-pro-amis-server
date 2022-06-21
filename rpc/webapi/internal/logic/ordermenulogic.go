package logic

import (
	"context"

	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderMenuLogic {
	return &OrderMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderMenuLogic) OrderMenu(in *webapi.OrderMenuReq) (*webapi.OrderMenuReply, error) {
	// todo: add your logic here and delete this line

	return &webapi.OrderMenuReply{}, nil
}
