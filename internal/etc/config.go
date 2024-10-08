package etc

import (
	"github.com/kataras/iris/v12"
	"time"
)

var Config Configuration

var TokenTTL = time.Hour * 24

var AppProxy *iris.Application
