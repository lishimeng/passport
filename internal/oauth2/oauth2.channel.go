package oauth2

import (
	"github.com/lishimeng/app-starter/midware/oauth/wx/miniwx"
	"github.com/lishimeng/go-log"
	"github.com/pkg/errors"
	"sync"
)

var wxMini miniwx.Oauth2Handle

var wxMiniOnce sync.Once

func GetWxMini() miniwx.Oauth2Handle {
	wxMiniOnce.Do(func() {
		c, err := getWxMiniConfig()
		if err != nil {
			log.Debug(errors.Wrapf(err, "wx_mini init failed"))
			return
		}
		wxMini = miniwx.New(miniwx.WithAuth(c.AppId, c.AppSecret))
	})
	return wxMini
}
