package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/internal/midware"
)

func Route(root iris.Party) {
	root.Get("/getUserInfo", midware.WithAuth(GetUserInfo)...)
	root.Post("/bindUser", midware.WithAuth(BindUser)...) //第三方关联绑定用户
}
