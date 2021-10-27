package icon

import (
	"runtime"

	"github.com/redhat-developer/app-services-cli/pkg/color"
)

const (
	ErrorSymbol     = "\u274c"
	checkMarkSymbol = "\u2714\ufe0f"
	infoSymbol      = "\u2139"
)

// Emoji accepts two arguments, emoji sequence code, and fallback string, for the cases when emoji isn't supported.
// If the running program's operating system target isn't Windows, then a function returns emoji, in other case fall back string.
func Emoji(emoji string, fallback string) string {
	if runtime.GOOS != "windows" {
		return emoji
	}
	return fallback
}

// ErrorPrefix returns cross mark emoji prefix or default "Error:"
func ErrorPrefix() string {
	return Emoji(ErrorSymbol, "Error:")
}

// SuccessPrefix returns check mark emoji prefix
func SuccessPrefix() string {
	emoji := rightPadPrefixIcon(checkMarkSymbol)

	return color.Success(emoji)
}

// InfoPrefix returns an emoji indicating an info message
func InfoPrefix() string {
	emoji := rightPadPrefixIcon(infoSymbol)
	return color.Info(emoji)
}

// Add a space after a prefix icon
func rightPadPrefixIcon(emojiCode string) string {
	fallback := ""
	emoji := Emoji(emojiCode, fallback)
	if emoji != fallback {
		emoji += " "
	}
	return emoji
}
