package page

import "github.com/kataras/iris/v12"

type IndexModel struct {
	Model
	Text string
}

func index(ctx iris.Context) {

	var err error
	var data IndexModel

	data.Title = "index.go设置title"
	data.Text = "这个页面是通过程序动态获取的"

	ctx.ViewLayout("layout/main")
	err = ctx.View("index.html", data)

	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}
