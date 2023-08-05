package page

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/midware/template"
	"github.com/lishimeng/passport/cmd/passport/static"
)

func application(app *iris.Application) {
	//if etc.Config.Web.Cache > 0 { // 设置了页面cache
	//	app.Use(iris.Cache304(time.Hour * time.Duration(etc.Config.Web.Cache)))
	//}

	engine := iris.HTML(static.Static, ".html")
	template.Init(engine)
	app.RegisterView(engine)
}

func Route(p *iris.Application) {

	application(p)
	// 这里不要用prefix为api的路径
	p.Get("/", index)
	p.Get("/login", login)
}
