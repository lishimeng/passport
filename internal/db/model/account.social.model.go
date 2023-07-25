package model

import "github.com/lishimeng/app-starter"

// SocialAccount 第三方账号关联
type SocialAccount struct {
	app.Pk
	AccountId       int            `orm:"column(account_id);null"`
	SocialAccountId string         `orm:"column(social_account_id);null"` // 第三方账号的key
	SocialGroupId   string         `orm:"column(social_group_id);null"`   // 对应微信union id,如没有可置空
	Category        SocialCategory `orm:"column(category);null"`
	app.TableChangeInfo
}

type SocialCategory string

const (
	SocialWeChat      SocialCategory = "wechat"      // 微信
	WeChatMiniProgram SocialCategory = "mini_wx"     // 微信小程序
	WeChatOfficial    SocialCategory = "official_wx" // 微信公众号
	WeChatQiye        SocialCategory = "qiye_wx"     // 企业微信
)
