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
	info.Status = 1
	_, err = app.GetOrm().Context.Insert(&info)
	return
}

func GetSocialAccountByAccountId(id int) (info model.SocialAccount, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.SocialAccount)).Filter("AccountId", id).One(&info)
	return
}

func GetSocialAccountById(socialAccountId, category string, uid int) (info model.SocialAccount, err error) {
	cond := orm.NewCondition()
	cond = cond.And("SocialAccountId", socialAccountId)
	if len(category) > 0 {
		cond = cond.And("Category", category)
	}
	if uid > 0 {
		cond = cond.And("AccountId", uid)
	}
	err = app.GetOrm().Context.QueryTable(new(model.SocialAccount)).SetCond(cond).One(&info)
	return
}

func GetTenantById(id int) (info model.Tenant, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Tenant)).Filter("Id", id).One(&info)
	return
}

func GetTenantAccountByUid(uid int) (info model.TenantAccount, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.TenantAccount)).Filter("Uid", uid).One(&info)
	return
}

func UpAccount(ori model.Account, cols ...string) (info model.Account, err error) {
	_, err = app.GetOrm().Context.Update(&ori, cols...)
	info = ori
	return
}
