package notify

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/sdk"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/etc"
)

// TODO 注册验证码(邮箱)
func SighUpSendMail(code string, to string) (resp sdk.Response, err error) {
	client := sdk.New(sdk.WithHost(etc.Config.Notify.Host), sdk.WithAuth(etc.Config.Notify.AppKey, etc.Config.Notify.Secret))
	template, err := GetSighUpEmailTemplate()
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

// 获取注册邮箱验证码模版
func GetSighUpEmailTemplate() (info model.Notify, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.Notify))
	cond := orm.NewCondition()
	cond = cond.And("Status__exact", 1)
	cond = cond.And("Category__exact", model.EmailSighup)
	err = qs.SetCond(cond).One(&info)
	return
}
