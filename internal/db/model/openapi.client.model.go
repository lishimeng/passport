package model

type OpenClient struct {
	Pk
	AppId    string `orm:"column(app_id);unique"`
	Secret   string `orm:"column(secret)"`
	Domain   string `orm:"column(domain)"`
	UserId   string `orm:"column(user_id);null"`
	Password string `orm:"column(password);null"`
	TableChangeInfo
}

func (t OpenClient) GetID() string {
	return t.AppId
}

func (t OpenClient) GetSecret() string {
	return t.Secret
}

func (t OpenClient) GetDomain() string {
	return t.Domain
}

func (t OpenClient) GetUserID() string {
	return t.AppId
}

func (t OpenClient) VerifyPassword(psw string) bool {
	return t.Password == psw
}
