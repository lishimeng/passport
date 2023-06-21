package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

type LoginReq struct {
	Password string `json:"password,omitempty"`
	UserName string `json:"userName,omitempty"`
	Code     string `json:"code,omitempty"`
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
	//todo 验证码校验
	if req.Code == "000000" {

	} else {

	}
	info, err := AccountLogin(req.UserName, req.Password)
	log.Debug("info:%s", info)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = 200
	resp.Uid = info.Id
	resp.Token = "11111111eeeeeeeeeeeeeeeeeeeeeee"
	tool.ResponseJSON(ctx, resp)
}
