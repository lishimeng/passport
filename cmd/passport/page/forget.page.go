package page

import (
	"github.com/kataras/iris/v12"
)

type forgetModel struct {
	Model
	Path string
}

func forget(ctx iris.Context) {
	var err error
	var data forgetModel
	path := ctx.URLParam("path")
	data.Title = "passport"
	data.Path = checkParams(path)
	ctx.ViewLayout("layout/main")
	err = ctx.View("forget.html", data)
	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}
