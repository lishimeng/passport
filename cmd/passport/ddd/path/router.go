package path

import (
	"github.com/lishimeng/app-starter/server"
)

func Router(root server.Router) {
	root.Get("", GetPathInfo) //登录
}
