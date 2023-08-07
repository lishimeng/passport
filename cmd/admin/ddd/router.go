package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/cmd/admin/ddd/account"
	"github.com/lishimeng/passport/cmd/admin/ddd/tenant"
)

func Route(root *iris.Application) {

	account.Route(root.Party("/account"))

	tenant.Route(root.Party("/tenant"))
}
