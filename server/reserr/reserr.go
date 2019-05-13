package reserr

// Error represents a RES error
type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

// RESError converts an error to an *Error. If it isn't of type *Error already, it will become a system.internalError.
func RESError(err error) *Error {
	rerr, ok := err.(*Error)
	if !ok {
		rerr = InternalError(err)
	}
	return rerr
}

// InternalError converts an error to an *Error with the code system.internalError.
func InternalError(err error) *Error {
	return &Error{Code: CodeInternalError, Message: "Internal error: " + err.Error()}
}

// Pre-defined RES error codes
const (
	CodeAccessDenied   = "system.accessDenied"
	CodeInternalError  = "system.internalError"
	CodeInvalidParams  = "system.invalidParams"
	CodeMethodNotFound = "system.methodNotFound"
	CodeNoSubscription = "system.noSubscription"
	CodeNotFound       = "system.notFound"
	CodeTimeout        = "system.timeout"
	CodeInvalidRequest = "system.invalidRequest"
	// HTTP only error codes
	CodeBadRequest         = "system.badRequest"
	CodeMethodNotAllowed   = "system.methodNotAllowed"
	CodeServiceUnavailable = "system.serviceUnavailable"
)

// Pre-defined RES errors
// https://github.com/resgateio/resgate/blob/master/docs/res-service-protocol.md#pre-defined-errors
// https://github.com/resgateio/resgate/blob/master/docs/res-client-protocol.md#pre-defined-errors
var (
	ErrAccessDenied   = &Error{Code: CodeAccessDenied, Message: "Access denied"}
	ErrDisposing      = &Error{Code: CodeInternalError, Message: "Internal error: disposing connection"}
	ErrInternalError  = &Error{Code: CodeInternalError, Message: "Internal error"}
	ErrInvalidParams  = &Error{Code: CodeInvalidParams, Message: "Invalid parameters"}
	ErrMethodNotFound = &Error{Code: CodeMethodNotFound, Message: "Method not found"}
	ErrNoSubscription = &Error{Code: CodeNoSubscription, Message: "No subscription"}
	ErrNotFound       = &Error{Code: CodeNotFound, Message: "Not found"}
	ErrTimeout        = &Error{Code: CodeTimeout, Message: "Request timeout"}
	ErrInvalidRequest = &Error{Code: CodeInvalidRequest, Message: "Invalid request"}
	// HTTP only errors
	ErrBadRequest         = &Error{Code: CodeBadRequest, Message: "Bad request"}
	ErrMethodNotAllowed   = &Error{Code: CodeMethodNotAllowed, Message: "Method not allowed"}
	ErrServiceUnavailable = &Error{Code: CodeServiceUnavailable, Message: "Service unavailable"}
)
