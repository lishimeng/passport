package main

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/cache"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/setup"
	"github.com/lishimeng/passport/cmd/passport/static"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/etc"
	"net/http"
	"time"
)
import _ "github.com/lib/pq"

func main() {
	orm.Debug = true

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error

		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		//配置过期时间
		if etc.Config.Ttl.Time <= 0 {
			etc.Config.Ttl.Time = 1
		}
		var day = time.Hour * 24 // 一天
		etc.TokenTTL = day * time.Duration(etc.Config.Ttl.Time)

		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}

		if etc.Config.Token.Enable {
			issuer := etc.Config.Token.Issuer
			tokenKey := []byte(etc.Config.Token.Key)
			builder = builder.EnableTokenValidator(func(inject app.TokenValidatorInjectFunc) {
				provider := token.NewJwtProvider(token.WithIssuer(issuer),
					token.WithKey(tokenKey, tokenKey), // hs256的秘钥必须是[]byte
					token.WithAlg("HS256"),
					token.WithDefaultTTL(etc.TokenTTL),
				)
				storage := token.NewLocalStorage(provider)
				factory.Add(provider)
				inject(storage)
			})
		}
		redisOpts := cache.RedisOptions{
			Addr:     etc.Config.Redis.Addr,
			Password: etc.Config.Redis.Password,
		}
		cacheOpts := cache.Options{
			MaxSize: 10000,
			Ttl:     etc.TokenTTL,
		}
		builder.EnableDatabase(dbConfig.Build(),
			model.Tables()...).
			SetWebLogLevel("DEBUG").
			EnableStaticWeb(func() http.FileSystem {
				return http.FS(static.Static)
			}).
			EnableCache(redisOpts, cacheOpts).
			EnableWeb(etc.Config.Web.Listen, setup.Application).
			PrintVersion()
		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
