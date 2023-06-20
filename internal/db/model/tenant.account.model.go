package model

import "github.com/lishimeng/app-starter"

// TenantAccount 账号--组织关系
type TenantAccount struct {
	app.TenantPk
	Uid int
	app.TableChangeInfo
}
