package account

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/passport/internal/db/model"
)

func createAccountSvc(name, phone, email string) (user model.Account, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) error {
		if len(name) > 0 {
			user.Name = name
		}
		if len(phone) > 0 {
			user.Mobile = phone
		}
		if len(email) > 0 {
			user.Name = email
		}
		user.Code = tool.GetUUIDString()
		// TODO
		_, e := ctx.Context.Insert(&user)
		if e != nil {
			return e
		}
		return nil
	})
	return
}
