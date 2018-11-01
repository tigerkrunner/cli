package wrapper

import "code.cloudfoundry.org/cli/api/router"

// ErrorWrapper is the wrapper that converts responses with 4xx and 5xx status
// codes to an error.
type ErrorWrapper struct {
	connection router.Connection
}

func NewErrorWrapper() *ErrorWrapper {
	return new(ErrorWrapper)
}
