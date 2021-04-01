package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type constraintError struct {
	ErrorID string
	Message string
}

func (err *constraintError) Error() string {
	return fmt.Sprintf("%s: %s", err.ErrorID, err.Message)
}

func InvalidFieldRangeLengthError(field string, min uint8, max uint) *constraintError {
	return &constraintError{
		ErrorID: "invalid-field-range-length-error",
		Message: fmt.Sprintf("'%s' must have between %d to %d characters.", field, min, max),
	}
}

func InvalidFieldMaxLengthError(field string, max uint) *constraintError {
	return &constraintError{
		ErrorID: "invalid-field-max-length-error",
		Message: fmt.Sprintf("'%s' must have a maximum of %d characters.", field, max),
	}
}

func ServiceNotFoundError(cause error) error {
	return errors.WithMessage(cause, "xxxx")
}

func ServiceAlreadyExistsError(name string) error {
	message := fmt.Sprintf("A Service with name '%s' already exists.", name)
	return errors.New(message)
}

//TODO The use cases errors should wrap with message the entities and adapters errors
func ConstraintError(cause error) error {
	return errors.Wrapf(cause, "Fail on validate business rules constraints.")
}
