/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-22 15:32:43
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-22 15:56:23
 * @FilePath     : \antd-pro-amis-server\rpc\webapi\internal\logic\sortmenulogic.go
 */
package logic

import (
	"context"
	"sort"

	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/webapi/internal/svc"
	"antd-pro-amis-server/rpc/webapi/webapi"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSortMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SortMenuLogic {
	return &SortMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func GetSortTree(ids *[]int32, sorts *[]int32, Rows []*webapi.SortMenuReq_Rows) {
	for _, v := range Rows {
		*ids = append(*ids, v.Id)
		*sorts = append(*sorts, v.Sort)
		if v.Children != nil {
			GetSortTree(ids, sorts, v.Children)
		}
	}
}

func (l *SortMenuLogic) SortMenu(in *webapi.SortMenuReq) (*webapi.SortMenuReply, error) {
	ids := []int32{}
	sorts := []int32{}
	GetSortTree(&ids, &sorts, in.Rows)

	sort.SliceStable(sorts, func(i, j int) bool {
		return sorts[i] < sorts[j] //从小到大排列
	})

	for i, v := range ids {
		if err := l.svcCtx.Db.Menu.Update().
			Where(menu.ID(int(v))).
			SetSort(int64(sorts[i])).
			Exec(context.Background()); err != nil {
			return nil, err
		}
	}
	return &webapi.SortMenuReply{Msg: "success"}, nil
}
