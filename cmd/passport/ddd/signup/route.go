package signup

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Post("/", register)
	root.Post("/phoneRegister", phoneRegister)
	root.Post("/emailRegister", emailRegister)
}
