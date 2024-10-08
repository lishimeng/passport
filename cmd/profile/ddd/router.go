package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/cms"
	"github.com/lishimeng/passport/cmd/profile/ddd/theme"
	"github.com/lishimeng/passport/cmd/profile/ddd/user"
)

func Route(root server.Router) {
	p := root.Party("/api")
	user.Route(p.Party("/user"))
	theme.Route(p.Party("/theme"))
	cms.Router(root)
}
