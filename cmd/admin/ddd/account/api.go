package account

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
)

type CreateAccountReq struct {
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

type CreateAccountResp struct {
	app.Response
	Id          int    `json:"id,omitempty"`
	AccountCode string `json:"accountCode,omitempty"`
	Name        string `json:"name,omitempty"`
}

// 只创建账号,不设置密码
func create(ctx iris.Context) {

	var req CreateAccountReq
	var err error
	var resp CreateAccountResp

	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Name) == 0 && len(req.Mobile) == 0 && len(req.Email) == 0 {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	user, err := createAccountSvc(req.Name, req.Mobile, req.Email)

	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Id = user.IdType
	resp.AccountCode = user.Code
	resp.Name = user.Name
	tool.ResponseJSON(ctx, resp)
}

func changePasswd(ctx iris.Context) {

}

func remove(ctx iris.Context) {

}
