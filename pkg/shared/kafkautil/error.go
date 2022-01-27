package kafkautil

import (
	"fmt"
)

// TODO - this methods should not be deifned here - or even defined this way.
type Error struct {
	Err error
}

// TODO - this methods should not be deifned here - or even defined this way.
func (e *Error) Error() string {
	return fmt.Sprint(e.Err)
}

// TODO - this methods should not be deifned here - or even defined this way.
func (e *Error) Unwrap() error {
	return e.Err
}

// TODO - we are mixing definitions of errors here and inline in many different occurences
func NotFoundByIDError(id string) error {
	NotFoundByIDErr := fmt.Errorf(`Kafka instance with ID "%v" not found`, id)
	return NotFoundByIDErr
}

func NotFoundByNameError(name string) error {
	NotFoundByNameErr := fmt.Errorf(`Kafka instance "%v" not found`, name)
	return NotFoundByNameErr
}

func InvalidSearchValueError(v string) error {
	IllegalSearchValueError := fmt.Errorf(`
	illegal search value "%v", search input must satisfy the following conditions:

  - must be of 1 or more characters
  - must only consist of alphanumeric characters, '-', '_' and '%%'
	`, v)

	return IllegalSearchValueError
}

func InvalidNameError(v string) error {
	InvalidNameErr := fmt.Errorf(`invalid Kafka instance name "%v". Valid names must satisfy the following conditions:

  - must be between 1 and 32 characters
  - must only consist of lower case, alphanumeric characters and '-'
  - must start with an alphabetic character
  - must end with an alphanumeric character
	`, v)
	return InvalidNameErr
}
