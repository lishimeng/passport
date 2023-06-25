package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/db/model"
)

type UserResp struct {
	app.Response
	Item model.Account `json:"item"`
}

func GetUserInfo(ctx iris.Context) {
	var resp UserResp
	//token := ctx.URLParam("token")
	id := ctx.Values().Get(auth.UidKey)
	log.Debug("id:%s", id)
	account, err := GetUserInfoById(id.(int))
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Item = account
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
