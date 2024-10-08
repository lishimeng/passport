package passwd

import (
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/x/digest"
)

// 生成密码
func Generate(plaintext string, ds model.Account) (r string) {
	r = digest.Generate(plaintext, ds.CreateTime.UnixNano())
	return
}

// 校验密码
func Verify(plaintext string, ds model.Account) (r bool) {
	r = digest.Verify(plaintext, ds.Password, ds.CreateTime.UnixNano())
	return
}
