/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-16 15:55:15
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-20 13:53:35
 * @FilePath     : \antd-pro-amis-server\server\internal\logic\getmenulogic.go
 */
package logic

import (
	"context"

	"antd-pro-amis-server/rpc/webapi/webapi"
	"antd-pro-amis-server/server/internal/svc"
	"antd-pro-amis-server/server/internal/types"

	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuLogic) GetMenu(req *types.GetMenuRequest) (resp []byte, err error) {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return []byte{}, err
	}

	bpReq := webapi.GetMenuReq{}
	if err := copier.Copy(&bpReq, req); err != nil {
		return []byte{}, err
	}

	if req.Id == nil {
		bpReq.Id = -1
	} else {
		bpReq.Id = *req.Id
	}

	if req.ParentId == nil {
		bpReq.ParentId = -1
	} else {
		bpReq.ParentId = *req.ParentId
	}

	res, err := l.svcCtx.Webapi.GetMenu(l.ctx, &bpReq)
	if err != nil {
		return []byte{}, err
	}

	res_json, err := json.Marshal(res)
	if err != nil {
		return []byte{}, err
	}

	return res_json, nil
}
