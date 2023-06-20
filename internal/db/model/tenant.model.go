package model

import "github.com/lishimeng/app-starter"

// Tenant 组织
type Tenant struct {
	app.Pk
	Name string `orm:"column(name);null"`
	app.TableChangeInfo
}
