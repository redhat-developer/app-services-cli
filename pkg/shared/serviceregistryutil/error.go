package serviceregistryutil

import (
	"fmt"
)

func NotFoundByIDError(id string) error {
	NotFoundByIDErr := fmt.Errorf(`Registry instance with ID "%v" not found`, id)
	return NotFoundByIDErr
}

func NotFoundByNameError(name string) error {
	NotFoundByNameErr := fmt.Errorf(`Service Registry instance "%v" not found`, name)
	return NotFoundByNameErr
}
