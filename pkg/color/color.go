// Color package is for printing a uniform set of colors for the CLI
package color

import (
	"github.com/fatih/color"
)

// CodeSnippet returns a colored string for code and command snippets
func CodeSnippet(format string) string {
	return color.HiMagentaString(format)
}

// Info returns a colored string for information messages
func Info(format string) string {
	return color.HiCyanString(format)
}

// Success returns a colored string for success messages
func Success(format string) string {
	return color.HiGreenString(format)
}

// Error returns a colored string for error messages
func Error(format string) string {
	return color.HiRedString(format)
}
