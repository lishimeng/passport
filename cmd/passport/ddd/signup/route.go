package signup

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Post("/", register)
	root.Post("/phoneRegister", phoneRegister)
	root.Post("/emailRegister", emailRegister)
}
