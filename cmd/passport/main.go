package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/api"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/etc"
	"github.com/lishimeng/owl/internal/setup"
	"github.com/lishimeng/owl/static"
	"net/http"
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
	time.Sleep(time.Second * 2)
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

		builder.EnableDatabase(dbConfig.Build(),
			model.Tables()...).
			//SetWebLogLevel("debug").
			PrintVersion().
			EnableWeb(etc.Config.Web.Listen, api.Route).
			EnableStaticWeb(func() http.FileSystem {
				return http.FS(static.Static)
			}).
			//ComponentBefore(setup.JobClearExpireTask).
			ComponentBefore(setup.BeforeStarted).
			ComponentAfter(setup.AfterStarted)

		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
