package connectorutil

import (
	"fmt"
)

func NotFoundByIDError(id string) error {
	NotFoundByIDErr := fmt.Errorf(`Coonector instance with ID "%v" not found`, id)
	return NotFoundByIDErr
}

func NotFoundByNameError(name string) error {
	NotFoundByNameErr := fmt.Errorf(`Coonector instance "%v" not found`, name)
	return NotFoundByNameErr
}
