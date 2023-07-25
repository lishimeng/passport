package oauth

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport/internal/db/model"
)

func getAppConfig(channel, appId string) (sc model.SocialConfig, err error) {

	if len(appId) > 0 {
		// 指定appId
		err = app.GetOrm().Context.
			QueryTable(new(model.SocialConfig)).
			Filter("AppId", appId).Filter("Status", app.Enable).One(&sc)
		return
	}
	err = app.GetOrm().Context.
		QueryTable(new(model.SocialConfig)).
		Filter("Channel", model.SocialCategory(channel)).Filter("Status", app.Enable).One(&sc)
	return
}
