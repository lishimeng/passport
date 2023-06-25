package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/passwd"
)

type LoginReq struct {
	Password  string `json:"password,omitempty"`
	UserName  string `json:"userName,omitempty"`
	Code      string `json:"code,omitempty"`
	LoginType string `json:"loginType,omitempty"`
}

type LoginResp struct {
	app.Response
	Token string `json:"token,omitempty"`
	Uid   int    `json:"uid,omitempty"`
}

func login(ctx iris.Context) {
	var resp LoginResp
	var req LoginReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	info, err := user.GetUserInfoByName(req.UserName)
	log.Debug("info:%s", info)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户名或密码错误"
		tool.ResponseJSON(ctx, resp)
		return
	}
	p := passwd.Verify(req.Password, info)
	log.Info("password:%s,%s", info.Password, p)
	if !p {
		resp.Code = tool.RespCodeError
		resp.Message = "用户名或密码错误"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Uid = info.Id
	log.Info("Uid:%s", info.Id)
	var provider token.JwtProvider
	err = factory.Get(&provider)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token 服务异常"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tokenPayload := token.JwtPayload{
		Org:    "org_1",
		Dept:   "dep_1",
		Uid:    info.Code,
		Client: req.LoginType,
	}
	tokenContent, err := provider.Gen(tokenPayload)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token生成失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Token = string(tokenContent)
	tool.ResponseJSON(ctx, resp)
}
