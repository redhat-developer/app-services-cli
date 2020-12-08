package kafka

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/MakeNowJust/heredoc"
)

var (
	validNameRegexp = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)
	errInvalidName  = errors.New(heredoc.Doc(`
	Invalid Kafka instance name. Valid names must satisfy the following conditions:

	- must be between 1 and 32 characters
	- must only consist of lower case, alphanumeric characters and '-'
	- must start with an alphabetic character
	- must end with an alphanumeric character
	`))
)

// ValidateName validates the proposed name of a Kafka instance
func ValidateName(name string) error {
	if len(name) < 1 || len(name) > 32 {
		return fmt.Errorf("Kafka instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return errInvalidName
}
