package user

import (
	"code-storm/common/result"
	"net/http"

	"code-storm/api/internal/logic/sys/user"
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserPageLogic(r.Context(), svcCtx)
		resp, err := l.UserPage(&req)
		result.HttpResult(r, w, resp, err)
	}
}
