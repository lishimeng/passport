package oauth

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/oauth/wx/miniwx"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/oauth2"
	"github.com/pkg/errors"
)

type AuthorizeTokenResp struct {
	app.Response

	AuthorizeToken string `json:"authorizeToken,omitempty"` // 用户的access token
	RefreshToken   string `json:"refreshToken,omitempty"`   // 用户的刷新Key
	OpenId         string `json:"openId,omitempty"`         // social端的用户ID
	UnionId        string `json:"unionId,omitempty"`        // social端的用户扩展ID
}

func authorizeCode(ctx iris.Context) {

	var err error
	var resp AuthorizeTokenResp
	channel := ctx.URLParam("channel")
	code := ctx.URLParam("code")
	if len(channel) == 0 || len(code) == 0 {
		log.Debug("channel:[%s] code:[%s]", channel, code)
		resp.Code = tool.RespCodeError

		tool.ResponseJSON(ctx, resp)
		return
	}
	switch model.SocialCategory(channel) {
	case model.WeChatMiniProgram:
		var wxToken miniwx.WxMiniLoginResp
		log.Debug("wx mini authorize code: %s", code)
		wxToken, err = getWxMiniToken(code)
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
		tool.ResponseJSON(ctx, resp)
		return
	}

	tool.ResponseJSON(ctx, resp)
}

func getWxMiniToken(code string) (resp miniwx.WxMiniLoginResp, err error) {
	handler := oauth2.GetWxMini()

	if handler != nil {
		resp, err = handler.AuthorizeCode(code)
		if err != nil {
			return
		}
	}
	return
}
