package signin

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {

	root.Post("/", login)
}
