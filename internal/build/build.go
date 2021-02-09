package build

import (
	"runtime/debug"
)

// Version is dynamically set by the toolchain or overridden by the Makefile.
var Version = "dev"
var Language = "en"

func init() {
	if Version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}
