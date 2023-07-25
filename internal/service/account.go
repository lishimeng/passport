package service

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/passport/internal/db/model"
)

func GenAccountCode() (code string) {
	code = tool.UUIDString()
	return
}

func GenTenantCode() (code string) {
	code = tool.UUIDString()
	return
}

// ResetPassword 直接修改密码, 如果不提供新密码,使用随机码创建密码
func ResetPassword(id int, newPasswd ...string) (uid int, code string, password string, err error) {

	psw := ""
	if len(newPasswd) > 0 {
		psw = newPasswd[0]
	} else {
		psw = tool.RandStr(12) // random password
	}
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) error {
		user := model.Account{}
		user.Id = id

		_, e := ctx.Context.Update(&user, "Password")
		if e != nil {
			return e
		}
		uid = user.Id
		code = user.Code
		password = psw
		return nil
	})
	return
}
