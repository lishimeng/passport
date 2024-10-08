package send

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Get("/signInSendCode", signInSendCodeGet) //登录
	root.Post("/signInSendCode", signInSendCodePost)
	root.Get("/signUpSendCode", signUpSendCodeGet) //注册
	root.Post("/signUpSendCode", signUpSendCodePost)
	root.Get("/bindSendCode", bindSendCodeGet) //绑定
	root.Post("/bindSendCode", bindSendCodePost)
}
