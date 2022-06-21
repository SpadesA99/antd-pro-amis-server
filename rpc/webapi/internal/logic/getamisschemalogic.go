package logic

import (
	"context"

	"antd-pro-amis-server/rpc/ent"
	"antd-pro-amis-server/rpc/ent/amisschema"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAmisSchemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAmisSchemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAmisSchemaLogic {
	return &GetAmisSchemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AmisSchema
func (l *GetAmisSchemaLogic) GetAmisSchema(in *webapi.GetAmisSchemaReq) (*webapi.GetAmisSchemaReply, error) {
	if schema, err := l.svcCtx.Db.AmisSchema.Query().Where(amisschema.Router(in.Router)).First(context.Background()); err != nil {
		if ent.IsNotFound(err) {
			return &webapi.GetAmisSchemaReply{Status: 200, Msg: "success", Data: &webapi.GetAmisSchemaReply_Data{Schema: `{"type":"page","body":[{"type":"code","language":"html","value":"<div>\r\n    <Show>\r\n        Hello FLY\r\n    </Show>\r\n</div>","id":"u:4c464d8f466a"}]}`}}, nil
		}

		return nil, err
	} else {
		return &webapi.GetAmisSchemaReply{Status: 200, Msg: "success", Data: &webapi.GetAmisSchemaReply_Data{Schema: schema.Schema}}, nil
	}
}
