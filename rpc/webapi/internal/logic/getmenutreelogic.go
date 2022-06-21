/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-20 16:19:29
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-21 13:55:20
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\getmenutreelogic.go
 */
package logic

import (
	"context"
	"strings"

	"antd-pro-amis-server/rpc/ent"
	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuTreeLogic {
	return &GetMenuTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type virtualMenuTree struct {
	ent.Menu
	Children []virtualMenuTree
}

func listToTree(raw_data []virtualMenuTree, MenuID int) []virtualMenuTree {
	var tree []virtualMenuTree
	for _, v := range raw_data {
		if int(v.ParentID) == MenuID {
			v.Children = listToTree(raw_data, int(v.ID))
			tree = append(tree, v)
		}
	}
	return tree
}

func parseTree(tree []virtualMenuTree, is_child bool, parent_path string) []*webapi.GetMenuTreeReply_Data {
	result := make([]*webapi.GetMenuTreeReply_Data, len(tree))

	for i := 0; i < len(tree); i++ {
		result[i] = &webapi.GetMenuTreeReply_Data{}
		if len(tree[i].Children) != 0 {
			//有子级说明这是一个父级菜单
			//例子
			/*
				path: '/menu',
				name: '菜单管理',
				roles: ['admin'],
				access: 'auth',
				wrappers: ['@/wrappers/auth'],
			*/
			result[i].Layout = true
			result[i].Path = tree[i].Path
			result[i].Name = tree[i].MenuName
			result[i].Routes = parseTree(tree[i].Children, true, tree[i].Path)
		} else {
			/*

			 */
			if !is_child {
				result[i].Access = "auth"
				//result[i].Wrappers = []string{"@/wrappers/auth"}
			}
			result[i].Roles = tree[i].Roles
			result[i].Path = tree[i].Path
			result[i].Name = tree[i].MenuName
			result[i].Layout = tree[i].Layout
			result[i].HideInMenu = tree[i].HideInMenu
			result[i].Icon = tree[i].Icon
			result[i].Component = "@/pages/Render"
			//result[i].Component = "./Render"
		}
	}
	if is_child {
		tmp := []*webapi.GetMenuTreeReply_Data{}
		//子级菜单需要添加一个重写路由在 0 位置
		res := append(tmp, &webapi.GetMenuTreeReply_Data{
			Path:     parent_path,
			Layout:   true,
			Redirect: result[0].Path,
		})
		res = append(res, result...)
		result = res
	}

	//不管是哪一组菜单都需要加上 404 页面
	//result = append(result, &webapi.GetMenuTreeReply_Data{Component: "@/pages/404"})
	return result
}

func (l *GetMenuTreeLogic) GetMenuTree(in *webapi.GetMenuTreeReq) (*webapi.GetMenuTreeReply, error) {
	if menus, err := l.svcCtx.Db.Menu.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueContains(menu.FieldRoles, strings.Split(in.Roles, ",")))
	}).Order(ent.Asc(menu.FieldSort)).All(context.Background()); err != nil {
		return nil, err
	} else {
		raw_data := make([]virtualMenuTree, len(menus))

		if err := copier.Copy(&raw_data, menus); err != nil {
			return nil, err
		}
		tree := listToTree(raw_data, 0)

		return &webapi.GetMenuTreeReply{Data: parseTree(tree, false, "")}, nil
	}

}
