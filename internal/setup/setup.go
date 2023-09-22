package setup

import (
	"context"
	"github.com/lishimeng/app-starter/cms"
)

func Setup(ctx context.Context) (err error) {
	cms.Init(cms.WithName("passport"), cms.WithDatabase())
	return
}
