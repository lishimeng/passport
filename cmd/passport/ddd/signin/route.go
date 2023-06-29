package signin

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Post("/login", login)
	root.Post("/codeLogin", codeLogin)
	root.Post("/socialLogin", openLogin)
}
