package openapi

import (
	"context"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-redis/redis/v8"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/etc"
	"sync"
)
import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
)

type Req struct {
	AccessToken  string
	RefreshToken string
	AppKey       string
	AppSecret    string
}

var (
	Srv  *server.Server
	once sync.Once
)

func Init(ctx context.Context) {

	manager := manage.NewDefaultManager()
	clientStore := NewClientStore()

	manager.MapClientStorage(clientStore)
	manager.MustTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: etc.Config.Redis.Addr,
	}), nil)

	once.Do(func() {
		Srv = server.NewDefaultServer(manager)
		Srv.SetAllowGetAccessRequest(true)
		Srv.SetClientInfoHandler(server.ClientFormHandler)
		Srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
			log.Info("Internal Error:%s", err.Error())
			return
		})
	})

}
