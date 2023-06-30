package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
)

func checkToken(ctx iris.Context) {
	var resp app.Response
	ui := ctx.Values().Get(auth.UserInfoKey)
	_, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
