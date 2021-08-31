package telemetry

import (
	"os"
	"os/user"
	"strings"
)

// Sanitizer replaces a PII data
const Sanitizer = "XXXX"

// SanitizeError sanitizes any PII(Personally Identifiable Information) from the error
func SanitizeError(err error) (errString string) {
	if err == nil {
		return ""
	}
	errString = err.Error()

	// Sanitize user information
	errString = sanitizeUserInfo(errString)

	// Sanitize file path
	errString = sanitizeFilePath(errString)

	return errString
}

// sanitizeUserInfo sanitizes username from the error string
func sanitizeUserInfo(errString string) string {
	user1, err1 := user.Current()
	if err1 != nil {
		return err1.Error()
	}
	errString = strings.ReplaceAll(errString, user1.Username, Sanitizer)
	return errString
}

// sanitizeFilePath sanitizes file paths from error string
func sanitizeFilePath(errString string) string {
	for _, str := range strings.Split(errString, " ") {
		if strings.Count(str, string(os.PathSeparator)) > 1 {
			errString = strings.ReplaceAll(errString, str, Sanitizer)
		}
	}
	return errString
}
