package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/cms"
	"github.com/lishimeng/passport/cmd/profile/ddd/theme"
	"github.com/lishimeng/passport/cmd/profile/ddd/user"
)

func Route(root *iris.Application) {

	p := root.Party("/api")
	user.Route(p.Party("/user"))
	theme.Route(p.Party("/theme"))
	cms.Router(root)
}
