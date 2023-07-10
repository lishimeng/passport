package model

func Tables() (t []interface{}) {
	t = append(t,
		new(OpenClient),
		new(Account),
		new(Tenant),
		new(TenantAccount),
		new(SocialAccount),
		new(Notify),
		new(ThemeConfig),
	)
	return
}
