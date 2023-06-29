package send

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Get("/sendCode", GetSendCode)
	root.Post("/sendCode", PostSendCode)
}
