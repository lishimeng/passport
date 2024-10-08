package page

import (
	"github.com/lishimeng/app-starter/server"
)

type logoutModel struct {
	Model
	Path string
}

func logout(ctx server.Context) {
	var err error
	var data logoutModel
	path := ctx.C.URLParam("path")
	data.Title = "passport"
	data.Path = checkParams(path)
	ctx.C.ViewLayout("layout/main")
	err = ctx.C.View("logout.html", data)
	if err != nil {
		_, _ = ctx.C.HTML("<h3>%s</h3>", err.Error())
	}
}
