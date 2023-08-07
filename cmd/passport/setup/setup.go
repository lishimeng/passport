package setup

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/cmd/passport/ddd"
	"github.com/lishimeng/passport/cmd/passport/page"
)

func Application(app *iris.Application) {

	//api := app.Party("/api")
	ddd.Route(app)

	page.Route(app)
}
