package setup

import (
	"context"
	"github.com/lishimeng/cms"
)

func Setup(ctx context.Context) (err error) {
	cms.Init(cms.WithName("passport-profile"), cms.WithDatabase())
	return
}
