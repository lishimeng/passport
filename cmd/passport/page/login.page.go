package page

import (
	"github.com/kataras/iris/v12"
)

type loginModel struct {
	Model
	Path string
}

func login(ctx iris.Context) {
	var err error
	var data loginModel
	path := ctx.URLParam("path")
	data.Title = "passport"
	data.Path = path
	ctx.ViewLayout("layout/main")
	err = ctx.View("login.html", data)
	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}

func phoneLogin(ctx iris.Context) {
	var err error
	var data loginModel
	path := ctx.URLParam("path")
	data.Title = "passport"
	data.Path = path
	ctx.ViewLayout("layout/main")
	err = ctx.View("phoneLogin.html", data)
	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}
