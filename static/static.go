package static

import (
	"embed"
)

//go:embed assets assets/css/* assets/js/* assets/svg/* favicon.ico index.html
var Static embed.FS
