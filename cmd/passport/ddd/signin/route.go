package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/internal/midware"
)

func Route(root iris.Party) {
	root.Post("/login", login)
	root.Post("/codeLogin", codeLogin)
	root.Post("/socialLogin", openLogin)
	root.Post("/checkToken", midware.WithAuth(checkToken)...)
}
