package oauth

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Get("/authorize_code", authorizeCode)
}
