package theme

import (
	"github.com/kataras/iris/v12"
)

func Route(root iris.Party) {
	root.Get("/", themeConfig)
}
