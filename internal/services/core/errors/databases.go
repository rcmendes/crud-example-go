package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type databaseError struct {
	ErrorID string
	Message string
	Cause   string
}

func (err *databaseError) Error() string {
	return fmt.Sprintf("[%s] %s. Cause:%v", err.ErrorID, err.Message, err.Cause)
}

func DBConnectionError(cause error) *databaseError {
	return &databaseError{
		ErrorID: "database-connection-error",
		Message: "Error when connecting into database",
		Cause:   fmt.Sprintf("Cause: %v", cause),
	}
}

func DBQueryError(cause error) *databaseError {
	return &databaseError{
		ErrorID: "database-query-error",
		Message: "Error while querying data from database",
		Cause:   fmt.Sprintf("Cause: %v", cause),
	}
}

func DatabaseError(cause error) error {
	return errors.Wrap(cause, "Database error.")
}
