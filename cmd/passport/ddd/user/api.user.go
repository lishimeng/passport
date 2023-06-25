package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

type UserResp struct {
	app.Response
	Item UserInfo `json:"item"`
}

type UserInfo struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Active int    `json:"active,omitempty"`
}

func GetUserInfo(ctx iris.Context) {
	var resp UserResp
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
	resp.Code = tool.RespCodeSuccess
	resp.Item.Id = account.Id
	resp.Item.Name = account.Name
	resp.Item.Email = account.Email
	resp.Item.Active = account.Active
	tool.ResponseJSON(ctx, resp)
	return
}

func tokenDecod(token string) (uid int, err error) {
	/*	var info *jwt.VerifiedToken
		info, err = token.VerifyToken(token)
		if err != nil {
			return
		}
		log.Debug("info :%s", info)
		log.Debug("info :%s", string(info.Payload))*/
	return
}
