package department

import (
	"code-storm/common/result"
	"net/http"

	"code-storm/api/internal/logic/sys/department"
	"code-storm/api/internal/svc"
	"code-storm/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddDepartmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddDepartmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := department.NewAddDepartmentLogic(r.Context(), svcCtx)
		resp, err := l.AddDepartment(&req)
		result.HttpResult(r, w, resp, err)
	}
}
