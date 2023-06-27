package model

import "github.com/lishimeng/app-starter"

// Notify 通知配置(模板:类型)
type Notify struct {
	app.Pk
	Template string         `orm:"column(template);null"`
	Category NotifyCategory `orm:"column(category);null"`
	app.TableChangeInfo
}

type NotifyCategory string

const (
	EmailSighup NotifyCategory = "email_sighup"  // email注册码
	SmsSighup   NotifyCategory = "sms_sighup"    // sms注册码
	EmailSighIn NotifyCategory = "email_sigh_in" // email登录码
	SmsSighIn   NotifyCategory = "sms_sigh_in"   // sms登录码
	EmailBind   NotifyCategory = "email_bind"    // email绑定
	SmsBind     NotifyCategory = "sms_bind"      // sms绑定
)
