package client

import (
	"code-storm/common/result"
	"net/http"

	"code-storm/api/internal/logic/sys/client"
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClientInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClientInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := client.NewClientInfoLogic(r.Context(), svcCtx)
		resp, err := l.ClientInfo(&req)
		/*if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}*/
		result.HttpResult(r, w, resp, err)
	}
}
