package account

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
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

type PasswordReq struct {
	Id       int    `json:"id,omitempty"`
	Password string `json:"password,omitempty"`
}

type PasswordResp struct {
	app.Response
}

func changePasswd(ctx iris.Context) {

	var err error
	var req PasswordReq
	var resp PasswordResp

	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
}

func remove(_ iris.Context) {

}

type GetInfoResp struct {
	app.Response
	AccountCode string
	Mobile      string
	Email       string
}

func info(ctx iris.Context) {
	var err error
	var resp GetInfoResp
	code := ctx.Params().Get("code")

	user, err := getAccountSvc(code)
	if err != nil {
		log.Debug(err)
		resp.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.AccountCode = user.Code
	resp.Email = user.Email
	resp.Mobile = user.Mobile

	tool.ResponseJSON(ctx, resp)
}
