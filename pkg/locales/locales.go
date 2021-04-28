package locales

import (
	"embed"
	"io/fs"
)

var (
	//go:embed files
	locales embed.FS
)

func FS() fs.FS {
	return locales
}
