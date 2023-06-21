package register

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
)

type RegisterReq struct {
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

func register(ctx iris.Context) {
	var resp app.Response
	var req RegisterReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	_, err = RegisterAccount(req.Mobile, req.Email, req.Password, req.Name)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "注册失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
