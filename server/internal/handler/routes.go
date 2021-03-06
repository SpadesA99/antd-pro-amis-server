// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"antd-pro-amis-server/server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add-menu-rols",
				Handler: AddMenuRolsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/del-menu-rols",
				Handler: DelMenuRolsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get-menu-rols",
				Handler: GetMenuRolsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get-schema",
				Handler: GetAmisSchemaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update-schema",
				Handler: UpdateAmisSchemaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add-menu",
				Handler: AddMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get-menu",
				Handler: GetMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/del-menu",
				Handler: DelMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update-menu",
				Handler: UpdateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/bluk-del-menu",
				Handler: BlukDelMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get-menu-tree",
				Handler: GetMenuTreeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sort-menu",
				Handler: SortMenuHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
