package signin

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

func loginView(ctx iris.Context) {

	redirectUri := ctx.URLParam("redirect_uri")
	if len(redirectUri) == 0 {
		ctx.StopWithStatus(http.StatusBadRequest)
	}

	_, _ = ctx.HTML("/#/login", "")
}
