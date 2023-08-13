package path

import "github.com/kataras/iris/v12"

func Router(root iris.Party) {
	root.Get("", GetPathInfo) //登录
}
