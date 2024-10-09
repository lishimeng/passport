package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/cms"
	"github.com/lishimeng/passport/cmd/profile/ddd/user"
)

func Route(root server.Router) {
	p := root.Path("/api")
	user.Route(p.Path("/user"))
	//theme.Route(p.Path("/theme"))
	cms.Router(p)
}
