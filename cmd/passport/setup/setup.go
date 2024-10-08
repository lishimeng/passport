package setup

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport/cmd/passport/ddd"
	"github.com/lishimeng/passport/cmd/passport/page"
)

func Application(app server.Router) {

	//api := app.Party("/api")
	ddd.Route(app)

	page.Route(app)
}
