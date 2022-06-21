package handler

import (
	"antd-pro-amis-server/server/internal/logic"
	"antd-pro-amis-server/server/internal/svc"
	"antd-pro-amis-server/server/internal/types"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetMenuRolsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type ErrorReply struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}
		var req types.GetMenuRolsRequest
		if err := httpx.Parse(r, &req); err != nil {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			res, _ := json.Marshal(&ErrorReply{Status: 400, Message: err.Error()})
			w.Write(res)
			return
		}

		l := logic.NewGetMenuRolsLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuRols(&req)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			res, _ := json.Marshal(&ErrorReply{Status: 400, Message: err.Error()})
			w.Write(res)
		} else {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if n, err := w.Write(resp); err != nil {
				// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
				// so it's ignored here.
				if err != http.ErrHandlerTimeout {
					logx.Errorf("write response failed, error: %s", err)
				}
			} else if n < len(resp) {
				logx.Errorf("actual bytes: %d, written bytes: %d", len(resp), n)
			}

		}
	}
}
