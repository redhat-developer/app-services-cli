package flagutil

// ValidateOutput checks if value v is a valid value for --output
func ValidateOutput(v string) error {
	isValid := IsValidInput(v, ValidOutputFormats...)

	if isValid {
		return nil
	}

	return InvalidValueError("output", v, ValidOutputFormats...)
}
