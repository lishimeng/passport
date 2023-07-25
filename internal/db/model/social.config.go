package model

import "github.com/lishimeng/app-starter"

type SocialConfig struct {
	app.Pk
	Channel   SocialCategory `orm:"column(channel)"`
	AppId     string         `orm:"column(appId)"`
	AppSecret string         `orm:"column(appSecret)"`
	app.TableChangeInfo
}
