package cmdutil

import (
	"fmt"
	"strconv"
)

func ConvertPageValueToInt32(s string) int32 {
	val, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 1
	}

	return int32(val)
}

func ConvertSizeValueToInt32(s string) int32 {
	val, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 10
	}

	return int32(val)
}

// StringSliceToListStringWithQuotes converts a string slice to a
// comma-separated list with each value in quotes.
// Example: "a", "b", "c"
func StringSliceToListStringWithQuotes(validOptions []string) string {
	var listF string
	for i, val := range validOptions {
		listF += fmt.Sprintf("\"%v\"", val)
		if i < len(validOptions)-1 {
			listF += ", "
		}
	}
	return listF
}
