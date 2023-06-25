package model

import "github.com/lishimeng/app-starter"

// Account 账号
type Account struct {
	app.Pk
	Code string `orm:"column(code)"`
	app.TableChangeInfo
	Profile
	IdCard
	AccountMobile
	AccountEmail
	AccountPassword
}

// AccountMobile MFA手机号
type AccountMobile struct {
	Mobile         string `orm:"column(mobile);null"`
	MobileVerified int    `orm:"column(mobile_verified);null"`
}

// AccountEmail MFA邮箱
type AccountEmail struct {
	Email         string `orm:"column(email);null"`
	EmailVerified int    `orm:"column(email_verified);null"`
}

// AccountPassword 密码配置
type AccountPassword struct {
	Password        string `orm:"column(password);null"`
	PasswordVersion int    `orm:"column(password_version);null"` // 比系统版本低就需要重置密码
}

// IdCard 证件配置
type IdCard struct {
	IdType     int    `orm:"column(id_type);null"`
	IdNumber   string `orm:"column(id_number);null"`
	IdVerified int    `orm:"column(id_verified);null"`
}

// Profile 用户信息
type Profile struct {
	Name   string `orm:"column(name);null"`   // 名称
	Avatar string `orm:"column(avatar);null"` // 头像
	Active int    `orm:"column(active);null"` // 激活(完成MFA其中一项都可以判定为激活)
}
