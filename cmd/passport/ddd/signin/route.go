package signin

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport/internal/midware"
)

func Route(root server.Router) {
	root.Post("/login", login)
	root.Post("/codeLogin", codeLogin)
	root.Post("/socialLogin", openLogin)
	root.Post("/checkToken", midware.WithAuth(checkToken)...)
	root.Post("/clearToken", midware.WithAuth(clearToken)...)
}
