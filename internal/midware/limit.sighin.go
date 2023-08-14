package midware

import (
	"fmt"
	"github.com/lishimeng/app-starter/factory"
	"time"
)

func CheckActionLimit(client string, action string) bool {

	const actionLimitTpl = "al_%s_%s"
	var key = fmt.Sprintf(actionLimitTpl, client, action)
	// TODO setNX替代
	c := factory.GetCache()

	exists := c.Exists(key)
	if exists {
		return false
	}
	_ = c.SetTTL(key, "1", time.Second) // TODO 1s
	return true
}
