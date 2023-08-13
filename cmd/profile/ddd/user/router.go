package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/internal/midware"
)

func Route(root iris.Party) {
	root.Get("/getUserInfo", midware.WithAuth(GetUserInfo)...)
	root.Post("/bindUser", midware.WithAuth(BindUser)...)   //第三方关联绑定用户
	root.Post("/bindPhone", midware.WithAuth(BindPhone)...) //换绑手机
	root.Post("/bindEmail", midware.WithAuth(BindEmail)...) //换绑邮箱

	root.Get("/bindSendCode", midware.WithAuth(bindSendCodeGet)...) //绑定
	root.Post("/changePassword", midware.WithAuth(changePassword)...)
}
