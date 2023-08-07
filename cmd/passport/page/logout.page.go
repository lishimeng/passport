package page

import (
	"github.com/kataras/iris/v12"
)

type logoutModel struct {
	Model
	Path string
}

func logout(ctx iris.Context) {
	var err error
	var data logoutModel
	path := ctx.URLParam("path")
	data.Title = "passport"
	data.Path = path
	ctx.ViewLayout("layout/main")
	err = ctx.View("logout.html", data)
	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}
