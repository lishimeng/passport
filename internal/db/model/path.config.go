package model

import "github.com/lishimeng/app-starter"

type PathConfig struct {
	app.Pk
	Path string `orm:"column(path)"`
}
