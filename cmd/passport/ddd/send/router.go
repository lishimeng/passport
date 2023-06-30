package send

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Get("/signInSendCode", signInSendCodeGet) //登录
	root.Post("/signInSendCode", signInSendCodePost)
	root.Get("/signUpSendCode", signUpSendCodeGet) //注册
	root.Post("/signUpSendCode", signUpSendCodePost)
	root.Get("/bindSendCode", bindSendCodeGet) //绑定
	root.Post("/bindSendCode", bindSendCodePost)
}
