package icon

import (
	"runtime"

	"github.com/redhat-developer/app-services-cli/pkg/color"
)

const (
	ErrorSymbol     = "\u274c"
	checkMarkSymbol = "\u2714\ufe0f"
	warningSymbol   = "\u26A0"
)

// Emoji accepts two arguments, emoji sequence code, and fallback string, for the cases when emoji isn't supported.
// If the running program's operating system target isn't Windows, then a function returns emoji, in other case fall back string.
func Emoji(emoji string, fallback string) string {
	if runtime.GOOS != "windows" {
		return emoji
	}
	return fallback
}

// SuccessPrefix returns check mark emoji prefix
func SuccessPrefix() string {
	return color.Success(Emoji(checkMarkSymbol, ""))
}

// ErrorPrefix returns cross mark emoji prefix or default "Error:"
func ErrorPrefix() string {
	return Emoji(ErrorSymbol, "Error:")
}

// Warning returns an emoji icon indicating a warning
// Ref: https://emojipedia.org/warning/
func Warning() string {
	return Emoji(warningSymbol, "")
}
