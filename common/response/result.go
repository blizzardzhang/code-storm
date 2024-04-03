package response

import "net/http"
import "github.com/zeromicro/go-zero/rest/httpx"

type Data struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(r *http.Request, w http.ResponseWriter, res any, err error) {
	if err != nil {
		data := Data{
			// 根据不同的错误码，返回不同的错误信息
			Code: http.StatusInternalServerError,
			Data: nil,
			Msg:  err.Error(),
		}
		httpx.WriteJson(w, http.StatusOK, data)
		return
	}
	data := Data{
		Code: http.StatusOK,
		Data: res,
		Msg:  "success",
	}
	httpx.WriteJson(w, http.StatusOK, data)
	return

}
