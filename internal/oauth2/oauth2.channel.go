package oauth2

import (
	"github.com/lishimeng/app-starter/midware/oauth/wx/miniwx"
	"sync"
)

var wxMini miniwx.Oauth2Handle

var wxMiniOnce sync.Once

func GetWxMini() miniwx.Oauth2Handle {
	wxMiniOnce.Do(func() {
		wxMini = miniwx.New()
	})
	return wxMini
}
