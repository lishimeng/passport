package signin

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/etc"
)

func AccountLogin(userName, password string) (info model.Account, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.Account))
	cond := orm.NewCondition()
	if len(userName) > 0 {
		cond = cond.AndCond(cond.Or("Name__exact", userName).Or("Mobile__exact", userName).Or("Email__exact", userName))
	}
	if len(password) > 0 {
		cond = cond.And("Password__exact", password)
	}
	err = qs.SetCond(cond).One(&info)
	if err != nil {
		return
	}
	return
}

func saveToken(tokenContent []byte) (err error) {
	key := token.Digest(tokenContent)
	err = app.GetCache().SetTTL(key, string(tokenContent), etc.TokenTTL)
	return
}
