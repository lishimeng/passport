package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

type checkTokenReq struct {
	Referre string `json:"referre,omitempty"`
}

func checkToken(ctx iris.Context) {
	var resp app.Response
	ui := ctx.Values().Get(auth.UserInfoKey)
	log.Info("token:%s", ui)
	_, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	var req checkTokenReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	referre := ctx.GetHeader("referre")
	log.Info("referre:%s,%S", referre, req.Referre)
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

// todo 退出登录 清理token
func clearToken(ctx iris.Context) {
	var resp app.Response
	tokenVal, _ := tool.GetAuth(ctx)
	log.Info("token:%s", tokenVal)
	key := token.Digest([]byte(tokenVal))
	log.Info("key:%s", key)
	if app.GetCache().Exists(key) {
		app.GetCache().Del(key)
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
