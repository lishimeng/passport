package oauth

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/oauth/wx/miniwx"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/pkg/errors"
)

type AuthorizeTokenResp struct {
	app.Response

	AuthorizeToken string `json:"authorizeToken,omitempty"` // 用户的access token
	RefreshToken   string `json:"refreshToken,omitempty"`   // 用户的刷新Key
	OpenId         string `json:"openId,omitempty"`         // social端的用户ID
	UnionId        string `json:"unionId,omitempty"`        // social端的用户扩展ID
}

func authorizeCode(ctx server.Context) {

	var err error
	var resp AuthorizeTokenResp
	channel := ctx.C.URLParam("channel")
	code := ctx.C.URLParam("code")
	appId := ctx.C.URLParam("appId")
	if len(code) == 0 {
		log.Debug("code:[%s]", code)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	if len(appId) > 0 {
		log.Debug("指定appId:%s", appId)
	}
	sc, err := getAppConfig(channel, appId)
	if err != nil {
		log.Debug(errors.Wrapf(err, "can't find app_config:[%s:%s]", channel, appId))
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	switch sc.Channel {
	case model.WeChatMiniProgram:
		var wxToken miniwx.WxMiniLoginResp
		log.Debug("wx mini authorize code: %s", code)

		wxToken, err = getWxMiniToken(sc.AppId, sc.AppSecret, code)
		if err == nil {
			resp.AuthorizeToken = wxToken.SessionKey
			resp.OpenId = wxToken.OpenId
			resp.UnionId = wxToken.UnionId
		}
	}

	if err != nil {
		log.Debug(errors.Wrapf(err, "get token failed"))
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func getWxMiniToken(appId, secret, code string) (resp miniwx.WxMiniLoginResp, err error) {
	handler := miniwx.New(miniwx.WithAuth(appId, secret))

	if handler != nil {
		resp, err = handler.AuthorizeCode(code)
		if err != nil {
			return
		}
	}
	return
}
