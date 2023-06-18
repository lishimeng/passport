package oauth

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-log"
	"net/http"
	"net/url"
)

const (
	loginTpl = `
<p>oauth2 login check</p>
<p id="code" style="display: none;">%s</p>
<script>

</script>
`
)

// login 验证用户是否登录
func login(ctx iris.Context) {
	redirectUri := ctx.URLParam("redirect_uri")
	if len(redirectUri) == 0 {
		ctx.StopWithStatus(http.StatusBadRequest)
	}
	_, err := ctx.HTML(loginTpl, "redirectUri")
	if err != nil {
		log.Info(err)
	}

	if isLogin(ctx) {
		ctx.Header("uid", "xxxxxx")
		ctx.Redirect(redirectUri)
	} else {
		location := fmt.Sprintf("/login?redirect_uri=%s", redirectUri)
		ctx.Redirect(url.QueryEscape(location))
	}

}

func isLogin(ctx iris.Context) bool {
	return false
}
