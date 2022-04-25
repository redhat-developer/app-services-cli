package cmdutil

import "time"

const (
	// The default indentation to use when printing data to stdout
	DefaultJSONIndent = "    "
	// DefaultPollTime is the default interval to wait when polling a network request
	DefaultPollTime = time.Millisecond * 5000
)
