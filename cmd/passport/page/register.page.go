package page

import (
	"github.com/kataras/iris/v12"
)

type registerModel struct {
	Model
	Path string
}

func register(ctx iris.Context) {
	var err error
	var data registerModel
	path := ctx.URLParam("path")
	data.Title = "passport"
	data.Path = checkParams(path)
	ctx.ViewLayout("layout/main")
	err = ctx.View("register.html", data)
	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}
