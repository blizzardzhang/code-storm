package app

import (
	"code-storm/common/result"
	"net/http"

	"code-storm/api/internal/logic/sys/app"
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateAppHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAppReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := app.NewUpdateAppLogic(r.Context(), svcCtx)
		resp, err := l.UpdateApp(&req)
		result.HttpResult(r, w, resp, err)
	}
}
