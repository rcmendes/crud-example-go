package errors

import "fmt"

type DatabaseError struct {
	ErrorID string
	Message string
	Cause   error
}

func (err *DatabaseError) Error() string {
	return fmt.Sprintf("[%s] %s. Cause:%v", err.ErrorID, err.Message, err.Cause)
}

func DBConnectionError(cause error) *DatabaseError {
	return &DatabaseError{
		ErrorID: "database-connection-error",
		Message: "Error when connecting into database",
		Cause:   cause,
	}
}

func DBQueryError(cause error) *DatabaseError {
	return &DatabaseError{
		ErrorID: "database-query-error",
		Message: "Error while querying data from database",
		Cause:   cause,
	}
}
