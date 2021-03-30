package exceptions

import "fmt"

type BusinessConstraintError struct {
	ErrorID string
	Message string
}

func (err *BusinessConstraintError) Error() string {
	return fmt.Sprintf("[%s] %s", err.ErrorID, err.Message)
}

func NewInvalidFieldRangeLengthError(field string, min uint8, max uint) *BusinessConstraintError {
	return &BusinessConstraintError{
		ErrorID: "invalid-field-range-length-error",
		Message: fmt.Sprintf("'%s' must have between %d to %d characters.", field, min, max),
	}
}

func NewInvalidFieldMaxLengthError(field string, max uint) *BusinessConstraintError {
	return &BusinessConstraintError{
		ErrorID: "invalid-field-max-length-error",
		Message: fmt.Sprintf("'%s' must have a maximum of %d characters.", field, max),
	}
}
