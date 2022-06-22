/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-16 15:51:48
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-22 15:27:10
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\getmenulogic.go
 */
package logic

import (
	"context"
	"strings"

	"antd-pro-amis-server/rpc/ent"
	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/utils"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type MvirtualMenuTree struct {
	ent.Menu
	Children []*webapi.GetMenuReply_Data_List
}

func MlistToTree(raw_data []MvirtualMenuTree, MenuID int) []*webapi.GetMenuReply_Data_List {
	var tree []*webapi.GetMenuReply_Data_List
	for _, v := range raw_data {
		if int(v.ParentID) == MenuID {
			v.Children = MlistToTree(raw_data, int(v.ID))
			tree = append(tree, &webapi.GetMenuReply_Data_List{
				Roles:      strings.Join(v.Roles, ","),
				Id:         int32(v.ID),
				MenuName:   v.MenuName,
				Path:       v.Path,
				ParentId:   v.ParentID,
				Icon:       v.Icon,
				Label:      v.Path,
				Sort:       int32(v.Sort),
				Value:      int32(v.ID),
				Status:     int32(utils.IfThen(v.Status, 1, 0)),
				Layout:     int32(utils.IfThen(v.Layout, 1, 0)),
				HideInMenu: int32(utils.IfThen(v.HideInMenu, 1, 0)),
				Children:   v.Children,
			})
		}
	}
	return tree
}

func (l *GetMenuLogic) GetMenu(in *webapi.GetMenuReq) (*webapi.GetMenuReply, error) {
	db := l.svcCtx.Db.Menu.Query()

	if in.Id != -1 {
		db = db.Where(menu.ID(int(in.Id)))
	}

	if in.ParentId != -1 {
		db = db.Where(menu.ParentID(in.ParentId))
	}

	if in.Status != -1 {
		db = db.Where(menu.Status(utils.IfThen(in.Status == 1, true, false)))
	}

	if in.Roles != "" {
		db = db.Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(menu.FieldRoles, in.Roles))
		})
	}

	//.Limit(int(in.PerPage)).Offset((int(in.Page) - 1) * int(in.PerPage))
	Menu, err := db.Order(ent.Asc(menu.FieldSort)).All(context.Background())
	if err != nil {
		return nil, err
	}

	raw_data := make([]MvirtualMenuTree, len(Menu))
	if err := copier.Copy(&raw_data, Menu); err != nil {
		return nil, err
	}

	tree := MlistToTree(raw_data, 0)

	return &webapi.GetMenuReply{
		Msg: "success",
		Data: &webapi.GetMenuReply_Data{
			Total: int32(len(tree)),
			Items: tree,
		}}, nil
}
