package notify

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport/internal/db/model"
)

// TODO 绑定(邮箱)
// 获取绑定邮箱验证码模版
func GetBindEmailTemplate() (info model.Notify, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.Notify))
	cond := orm.NewCondition()
	cond = cond.And("Status__exact", 1)
	cond = cond.And("Category__exact", model.EmailBind)
	err = qs.SetCond(cond).One(&info)
	return
}
