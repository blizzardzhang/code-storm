package client

import (
	"code-storm/common/response"
	"net/http"

	"code-storm/api/internal/logic/sys/client"
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClientAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddClientReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := client.NewClientAddLogic(r.Context(), svcCtx)
		resp, err := l.ClientAdd(&req)
		response.Result(r, w, resp, err)
	}
}
