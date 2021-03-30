package database

import "fmt"

type DatabaseError struct {
	ErrorID   string
	Message   string
	RootError error
}

func (err *DatabaseError) Error() string {
	return fmt.Sprintf("[%s] %s.", err.ErrorID, err.Message)
}

func NewDBConnectionError(rootError error) *DatabaseError {
	return &DatabaseError{
		ErrorID:   "database-connection-error",
		Message:   "Error when connecting into database.",
		RootError: rootError,
	}
}

func NewDBQueryError(rootError error) *DatabaseError {
	return &DatabaseError{
		ErrorID:   "database-query-error",
		Message:   fmt.Sprintf("Error when querying data from database. Error: %v", rootError),
		RootError: rootError,
	}
}
