/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-20 14:54:42
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-20 14:54:50
 * @FilePath     : \antd-pro-amis-server\server\internal\logic\blukdelmenulogic.go
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

type BlukDelMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlukDelMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlukDelMenuLogic {
	return &BlukDelMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlukDelMenuLogic) BlukDelMenu(req *types.BlukelMenuRequest) (resp []byte, err error) {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return []byte{}, err
	}

	bpReq := webapi.BulkDelMenuReq{}
	if err := copier.Copy(&bpReq, req); err != nil {
		return []byte{}, err
	}

	res, err := l.svcCtx.Webapi.BulkDelMenu(l.ctx, &bpReq)
	if err != nil {
		return []byte{}, err
	}

	res_json, err := json.Marshal(res)
	if err != nil {
		return []byte{}, err
	}

	return res_json, nil
}
