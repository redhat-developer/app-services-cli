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

// Success returns check mark emoji
func Success() string {
	return Emoji("\u2714\ufe0f", "")
}

// Error returns cross mark emoji
func Error() string {
	return Emoji("\u274c", "Error:")
}
