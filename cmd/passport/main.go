package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/passport/cmd/passport/ddd"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/etc"
	"time"
)
import _ "github.com/lib/pq"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 100)
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
				provider := token.NewJwtProvider(issuer,
					token.WithKey(tokenKey, tokenKey), // hs256的秘钥必须是[]byte
					token.WithAlg("HS256"),
				)
				storage := token.NewLocalStorage(provider)
				inject(storage)
			})
		}

		builder.EnableDatabase(dbConfig.Build(),
			model.Tables()...).
			//SetWebLogLevel("debug").
			PrintVersion().
			EnableWeb(etc.Config.Web.Listen, ddd.Route)
		/*.
		EnableStaticWeb(func() http.FileSystem {
			return http.FS(static.Static)
		})*/
		//ComponentBefore(setup.JobClearExpireTask).
		//ComponentBefore(setup.BeforeStarted).
		//ComponentAfter(setup.AfterStarted)

		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
