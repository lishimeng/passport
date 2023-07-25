package oauth2

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport/internal/db/model"
)

func getWxMiniConfig() (sc model.SocialConfig, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.SocialConfig)).
		Filter("Channel", model.WeChatMiniProgram).One(&sc)
	return
}
