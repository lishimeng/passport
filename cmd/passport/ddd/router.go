package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport/cmd/passport/ddd/oauth"
	"github.com/lishimeng/passport/cmd/passport/ddd/path"
	"github.com/lishimeng/passport/cmd/passport/ddd/send"
	"github.com/lishimeng/passport/cmd/passport/ddd/signin"
	"github.com/lishimeng/passport/cmd/passport/ddd/signup"
	"github.com/lishimeng/passport/cmd/passport/ddd/theme"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
)

func Route(root server.Router) {

	p := root.Path("/api")
	signin.Route(p.Path("/sign_in"))
	signup.Route(p.Path("/register"))
	user.Route(p.Path("/user"))
	send.Route(p.Path("/notify"))
	theme.Route(p.Path("/theme"))
	oauth.Route(p.Path("/oauth"))
	path.Router(p.Path("/path"))
}
