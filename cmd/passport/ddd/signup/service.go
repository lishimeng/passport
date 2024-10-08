package signup

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/x/util"
)

func RegisterAccount(mobile, email, password, name string) (info model.Account, err error) {
	if len(mobile) > 0 {
		info.Mobile = mobile
	}
	if len(email) > 0 {
		info.Email = email
	}
	if len(password) > 0 {
		info.Password = password
	}
	if len(name) > 0 {
		info.Name = name
	}
	info.Status = model.ActivateEnable
	info.Code = util.UUIDString() // 随机code 以后升级一下lib,改为tool.uuid
	_, err = app.GetOrm().Context.Insert(&info)
	return
}
