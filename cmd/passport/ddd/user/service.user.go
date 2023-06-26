package user

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport/internal/db/model"
)

func GetUserInfoById(id int) (info model.Account, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Account)).Filter("Id", id).One(&info)
	return
}

func GetUserInfoByName(name string) (info model.Account, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Account)).Filter("Name", name).One(&info)
	return
}

func GetUserInfoByCode(code string) (info model.Account, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Account)).Filter("Code", code).One(&info)
	return
}

func GetUserInfoByUserName(userName string) (info model.Account, err error) {
	cond := orm.NewCondition()
	cond = cond.Or("Name", userName)
	cond = cond.Or("Mobile", userName)
	cond = cond.Or("Email", userName)
	err = app.GetOrm().Context.QueryTable(new(model.Account)).SetCond(cond).One(&info)
	return
}

func GetUserInfoByThree(name, mobile, email string) (info model.Account, err error) {
	cond := orm.NewCondition()
	cond = cond.Or("Name", name)
	cond = cond.Or("Mobile", mobile)
	cond = cond.Or("Email", email)
	err = app.GetOrm().Context.QueryTable(new(model.Account)).SetCond(cond).One(&info)
	return
}

func InsertSocialAccount(socialAccount model.SocialAccount) (info model.SocialAccount, err error) {
	info = socialAccount
	_, err = app.GetOrm().Context.Insert(&info)
	return
}

func GetSocialAccountById(socialAccountId string) (info model.SocialAccount, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.SocialAccount)).Filter("SocialAccountId", socialAccountId).One(&info)
	return
}
