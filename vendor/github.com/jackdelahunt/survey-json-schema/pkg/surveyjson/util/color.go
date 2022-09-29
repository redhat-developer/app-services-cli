package util

import (
	"github.com/fatih/color"
)

// ColorInfo returns a new function that returns info-colorized (green) strings for the
// given arguments with fmt.Sprint().
var ColorInfo = color.New(color.FgGreen).SprintFunc()
