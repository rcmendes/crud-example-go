package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type businessConstraintError struct {
	ErrorID string
	Message string
}

func (err *businessConstraintError) Error() string {
	return fmt.Sprintf("[%s] %s", err.ErrorID, err.Message)
}

func InvalidFieldRangeLengthError(field string, min uint8, max uint) *businessConstraintError {
	return &businessConstraintError{
		ErrorID: "invalid-field-range-length-error",
		Message: fmt.Sprintf("'%s' must have between %d to %d characters.", field, min, max),
	}
}

func InvalidFieldMaxLengthError(field string, max uint) *businessConstraintError {
	return &businessConstraintError{
		ErrorID: "invalid-field-max-length-error",
		Message: fmt.Sprintf("'%s' must have a maximum of %d characters.", field, max),
	}
}

//TODO The use cases errors should wrap with message the entities and adapters errors
func ServiceNotFoundError(cause error) error {
	return errors.WithMessage(cause, "xxxx")
}

func BusinessConstraintError(cause error) error {
	return errors.WithMessage(cause, "xxxx")
}
