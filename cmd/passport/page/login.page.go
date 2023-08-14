package page

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
	"strings"
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
	data.Path = checkParams(path)
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
	data.Path = checkParams(path)
	ctx.ViewLayout("layout/main")
	err = ctx.View("phoneLogin.html", data)
	if err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
	}
}

func checkParams(path string) (p string) {
	log.Info("path: %s", path)
	if len(path) == 0 {
		return
	}
	if strings.Index(path, "?") >= 0 {
		np := path[0:strings.Index(path, "?")]
		log.Info("host: %s", np)
		params := path[strings.Index(path, "?")+1:]
		log.Info("params: %s", params)
		if strings.Index(params, "&") >= 0 {
			maps := strings.Split(params, "&")
			var s string
			for _, v := range maps {
				//log.Info("v:%s,%d", v, i)
				if strings.Index(v, "token") < 0 {
					if len(s) > 0 {
						s = v + "&" + s
					} else {
						s = v
					}
				}
				log.Info("s:%s", s)
			}
			p = np + "?" + s
		} else {
			if strings.Index(params, "token") < 0 {
				p = np + "?" + params
			} else {
				p = np
			}
		}
	} else {
		p = path
	}
	log.Info("path:%s", p)
	return
}
