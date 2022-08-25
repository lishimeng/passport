package openapi

import (
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/passport/internal/db/model"
)

func GetClientByAppKey(ctx persistence.OrmContext, appKey string) (ci model.OpenClient, err error) {
	return
}
