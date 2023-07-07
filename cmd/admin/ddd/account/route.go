package account

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/passport/internal/midware"
)

func Route(root iris.Party) {

	root.Get("/{code:string}", midware.WithAuth(info)...)
	root.Post("/", midware.WithAuth(create)...)
	root.Delete("/", midware.WithAuth(remove)...)
	root.Post("/{id:int}/password", midware.WithAuth(changePasswd)...)
}
