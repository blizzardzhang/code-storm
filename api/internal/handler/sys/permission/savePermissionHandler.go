package permission

import (
	"code-storm/common/result"
	"net/http"

	"code-storm/api/internal/logic/sys/permission"
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SavePermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SavePermissionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permission.NewSavePermissionLogic(r.Context(), svcCtx)
		resp, err := l.SavePermission(&req)
		result.HttpResult(r, w, resp, err)
	}
}
