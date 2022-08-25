package openapi

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
)

func AuthorizationCk(ctx iris.Context) {
	_, success := Authorization(ctx)
	if !success {
		var resp = app.Response{}
		resp.Code = 401
		return
	}
	ctx.Next()
}

func Authorization(ctx iris.Context) (claims *jwt.Claims, success bool) {

	return
}
