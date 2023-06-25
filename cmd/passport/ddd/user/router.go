package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/internal/midware"
)

func Route(root iris.Party) {
	root.Get("/getUserInfoByToken", midware.WithAuth(GetUserInfo)...)
}
