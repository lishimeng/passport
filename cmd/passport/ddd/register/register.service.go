package register

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport/internal/db/model"
)

func RegisterAccount(mobile, email, password, name string) (info model.Account, err error) {
	info.Mobile = mobile
	info.Email = email
	info.Password = password
	info.Name = name
	_, err = app.GetOrm().Context.Insert(&info)
	return
}
