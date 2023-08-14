package path

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/passport/internal/db/model"
)

type pathInfo struct {
	app.Response
	Path string `json:"path"`
}

func GetPathInfo(ctx iris.Context) {
	var resp pathInfo
	var info model.PathConfig
	err := app.GetOrm().Context.QueryTable(new(model.PathConfig)).One(&info)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "尚未配置Path！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Path = info.Path
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
