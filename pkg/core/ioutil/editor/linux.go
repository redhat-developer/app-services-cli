// nolint
//go:build !windows
// +build !windows

package editor

const (
	defaultShell     = "sh"
	shellCommandFlag = "-c"
	defaultEditor    = "vi"
)
