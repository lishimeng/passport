package static

import (
	"embed"
)

//go:embed static/* assets/* favicon.ico index.html
var Static embed.FS
