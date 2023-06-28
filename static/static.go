package static

import (
	"embed"
)

//go:embed assets/* favicon.ico index.html
var Static embed.FS
