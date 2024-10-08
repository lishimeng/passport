package theme

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Get("/", themeConfig)
}
