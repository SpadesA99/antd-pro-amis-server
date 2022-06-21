/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-20 16:42:40
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-20 16:42:52
 * @FilePath     : \antd-pro-amis-server\server\internal\logic\getmenutreelogic.go
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

type GetMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuTreeLogic {
	return &GetMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuTreeLogic) GetMenuTree(req *types.GetMenuTreeRequest) (resp []byte, err error) {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return []byte{}, err
	}

	bpReq := webapi.GetMenuTreeReq{}
	if err := copier.Copy(&bpReq, req); err != nil {
		return []byte{}, err
	}

	res, err := l.svcCtx.Webapi.GetMenuTree(l.ctx, &bpReq)
	if err != nil {
		return []byte{}, err
	}

	res_json, err := json.Marshal(res)
	if err != nil {
		return []byte{}, err
	}

	return res_json, nil
}
