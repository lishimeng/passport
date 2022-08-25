package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl/ddd/sender"
	"github.com/lishimeng/owl/ddd/um"
)

func Router(root iris.Party) {

	um.Route(root.Party("/messages"))
	sender.Route(root.Party("/sender"))
}
