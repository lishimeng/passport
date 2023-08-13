package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/passwd"
)
type PasswordReq struct {
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

func changePassword(ctx iris.Context) {
	var resp app.Response
	var req PasswordReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	account, err := user.GetUserInfoByUserName(req.Name)
	log.Debug("info:%s", account)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户不存在"
		tool.ResponseJSON(ctx, resp)
		return
	}
	p := passwd.Generate(req.Password, account)
	account.Password = p
	_, err = user.UpAccount(account, "password")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "更新密码失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
