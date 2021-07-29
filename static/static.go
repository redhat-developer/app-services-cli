package static

import (
	"embed"
	"io/fs"
)

//go:embed img/*
var images embed.FS

// ImagesFS returns the embedded images assets
func ImagesFS() fs.FS {
	return images
}
