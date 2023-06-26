package signup

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/passport/internal/db/model"
)

func RegisterAccount(mobile, email, password, name string) (info model.Account, err error) {
	info.Mobile = mobile
	info.Email = email
	info.Password = password
	info.Name = name
	info.Code = tool.GetRandomString(16) // 随机code 以后升级一下lib,改为tool.uuid
	_, err = app.GetOrm().Context.Insert(&info)
	return
}

func upPassword(ori model.Account, cols ...string) (info model.Account, err error) {
	_, err = app.GetOrm().Context.Update(&ori, cols...)
	info = ori
	return
}
