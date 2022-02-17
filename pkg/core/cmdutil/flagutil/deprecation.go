package flagutil

// DeprecateFlag provides a way to deprecate a flag by appending standard prefixes to the flag description.
func DeprecateFlag(flagDescription string) string {
	return "DEPRECATED: " + flagDescription
}
