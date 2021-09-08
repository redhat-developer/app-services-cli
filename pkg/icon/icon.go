package icon

import "runtime"

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
	return Emoji("\u2714\ufe0f", "")
}

// ErrorPrefix returns cross mark emoji prefix or default "Error:"
func ErrorPrefix() string {
	return Emoji("\u274c", "Error:")
}
