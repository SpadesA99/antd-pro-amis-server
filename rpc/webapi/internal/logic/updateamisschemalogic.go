/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-15 13:48:42
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-15 13:57:55
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\updateamisschemalogic.go
 */
package logic

import (
	"context"

	"antd-pro-amis-server/rpc/ent/amisschema"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAmisSchemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAmisSchemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAmisSchemaLogic {
	return &UpdateAmisSchemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAmisSchemaLogic) UpdateAmisSchema(in *webapi.UpdateAmisSchemaReq) (*webapi.UpdateAmisSchemaReply, error) {

	if b := l.svcCtx.Db.AmisSchema.Query().Where(amisschema.Router(in.Router)).ExistX(context.Background()); !b {

		if err := l.svcCtx.Db.AmisSchema.Create().SetRouter(in.Router).SetSchema(in.Schema).Exec(context.Background()); err != nil {
			return nil, err
		}

	} else {

		if err := l.svcCtx.Db.AmisSchema.Update().Where(amisschema.Router(in.Router)).SetSchema(in.Schema).Exec(context.Background()); err != nil {
			return nil, err
		}

	}

	return &webapi.UpdateAmisSchemaReply{Status: 200, Msg: "success"}, nil
}
