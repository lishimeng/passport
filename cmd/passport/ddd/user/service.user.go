package user

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport/internal/db/model"
)

func GetUserInfoById(id int) (info model.Account, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Account)).Filter("Id", id).One(&info)
	return
}

func GetUserInfoByCode(code string) (info model.Account, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Account)).Filter("Code", code).One(&info)
	return
}

func GetUserInfoByName(name string) (info model.Account, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Account)).Filter("Name", name).One(&info)
	return
}
