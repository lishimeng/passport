package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
)

func login(ctx iris.Context) {

	var resp app.Response
	resp.Code = 200
	tool.ResponseJSON(ctx, resp)
}
