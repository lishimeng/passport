package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/cmd/passport/ddd/send"
	"github.com/lishimeng/passport/cmd/passport/ddd/signin"
	"github.com/lishimeng/passport/cmd/passport/ddd/signup"
	"github.com/lishimeng/passport/cmd/passport/ddd/theme"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
)

func Route(root *iris.Application) {

	p := root.Party("/api")
	signin.Route(p.Party("/sign_in"))
	signup.Route(p.Party("/register"))
	user.Route(p.Party("/user"))
	send.Route(p.Party("/notify"))

	theme.Route(p.Party("/theme"))
}
