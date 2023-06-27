package notify

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/sdk"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/etc"
)

// TODO 绑定(邮箱)
func BindSendMail(code string, to string) (resp sdk.Response, err error) {
	client := sdk.NewClient(etc.Config.Notify.Host)
	template, err := GetBindEmailTemplate()
	if err != nil {
		return
	}
	params := make(map[string]string)
	params["verrificationCode"] = code
	var req = sdk.MailRequest{
		Template:      template.Template,
		CloudTemplate: false,
		TemplateParam: params,
		Title:         "验证码",
		Receiver:      to,
	}
	resp, err = client.SendMail(req)
	if err != nil {
		return
	}
	return
}

// 获取绑定邮箱验证码模版
func GetBindEmailTemplate() (info model.Notify, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.Notify))
	cond := orm.NewCondition()
	cond = cond.And("Status__exact", 1)
	cond = cond.And("Category__exact", model.EmailBind)
	err = qs.SetCond(cond).One(&info)
	return
}
