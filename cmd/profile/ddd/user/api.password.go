package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/passwd"
)
type PasswordReq struct {
	Password string `json:"password,omitempty"`
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
	ui := ctx.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("code:%s", jwt.Uid)
	account, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
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
