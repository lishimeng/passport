package etc

type Configuration struct {
	Db     db
	Web    web
	Redis  redis
	Token  token
	Notify notify
	Ttl    ttl
}

type web struct {
	Listen string
}

type redis struct {
	Addr     string
	Password string
	Db       int
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}

type token struct {
	Enable bool
	Issuer string
	Key    string
}

type notify struct {
	Host   string
	AppKey string
	Secret string
}

type ttl struct {
	Time     int64
	TimeType string
}
