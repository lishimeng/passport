package notify

import (
	"github.com/lishimeng/owl-messager/sdk"
	"github.com/lishimeng/passport/internal/etc"
)

// TODO 绑定(SMS)

func BindPhone(code string, to string) (resp sdk.Response, err error) {
	client := sdk.NewClient(etc.Config.Notify.Host)
	var req = sdk.SmsRequest{} // TODO
	resp, err = client.SendSms(req)
	return
}
