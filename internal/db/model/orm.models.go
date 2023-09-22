package model

import "github.com/lishimeng/app-starter/cms"

func Tables() (t []interface{}) {
	t = append(t,
		//new(OpenClient),
		new(Account),
		new(Tenant),
		new(TenantAccount),
		new(SocialAccount),
		new(SocialConfig),
		new(Notify),
		new(ThemeConfig),
		new(PathConfig),
		new(cms.SpaConfig),
	)
	return
}
