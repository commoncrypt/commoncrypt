package errors

import "fmt"

type Error struct {
	Code    string
	Message string
	Context string
	Status  int
}

var (
	InvalidValue  = Error{Code: "InvalidValue", Message: "An invalid value was provided to the server", Status: 400}
	AlreadyExists = Error{Code: "AlreadyExists", Message: "The resource already exists", Status: 409}
)

func (e Error) Error() string {
	if e.Context != "" {
		return fmt.Sprint(e.Message, ": ", e.Context)
	}

	return e.Message
}

func (e Error) WithContext(context string) Error {
	e.Context = context
	return e
}
