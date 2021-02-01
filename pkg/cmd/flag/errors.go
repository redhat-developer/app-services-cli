package flag

import "fmt"

func InvalidArgumentError(flag string, value string, err error) error {
	return fmt.Errorf(`invalid argument "%v" for "%v" flag: %w`, flag, value, err)
}
