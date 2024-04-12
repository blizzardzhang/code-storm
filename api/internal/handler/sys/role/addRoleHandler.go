package role

import (
	"code-storm/common/result"
	"net/http"

	"code-storm/api/internal/logic/sys/role"
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewSaveRoleLogic(r.Context(), svcCtx)
		resp, err := l.SaveRole(&req)
		result.HttpResult(r, w, resp, err)
	}
}
