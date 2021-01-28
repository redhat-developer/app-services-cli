package kafka

import (
	"fmt"
)

type Error struct {
	Reason string
}

func (e *Error) Error() string {
	return e.Reason
}

func ErrorNotFound(id string) *Error {
	return &Error{
		Reason: fmt.Sprintf("Kafka instance with ID '%v' not found", id),
	}
}
