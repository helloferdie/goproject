package liberror

type Base struct {
	Error     string
	ErrorVars map[string]interface{}
}

type Error struct {
	Type   string
	Errors []*Base
}

var (
	TypeRequest    = "common.error.request.default"
	TypeServer     = "common.error.server.default"
	TypeValidation = "common.error.validation.default"
)

// Error
var (
	// Request
	ErrForbidden    = "common.error.request.forbidden"
	ErrUnauthorized = "common.error.request.unauthorized"
	ErrNotFound     = "common.error.request.not_found.data"

	// Server
	ErrRepository     = "common.error.server.repository"
	ErrNotImplemented = "common.error.server.not_implemented"
)

func NewError(errorType string, list []*Base) *Error {
	return &Error{
		Type:   errorType,
		Errors: list,
	}
}

func NewErrorQuick(errorType string, errorMessage string) *Error {
	return NewError(errorType, []*Base{{
		Error:     errorMessage,
		ErrorVars: nil,
	}})
}

func NewServerError(errMsg string) *Error {
	return NewErrorQuick(TypeServer, errMsg)
}

func NewRequestError(errMsg string) *Error {
	return NewErrorQuick(TypeRequest, errMsg)
}

func NewErrNotImplemented() *Error {
	return NewServerError(ErrNotImplemented)
}

func NewErrRepository() *Error {
	return NewServerError(ErrRepository)
}

func NewErrNotFound(msg string) *Error {
	if msg == "" {
		msg = "common.error.request.not_found.data"
	}
	return NewErrorQuick("common.error.request.not_found.default", msg)
}

func NewErrValidation(fields ...*Base) *Error {
	return &Error{
		Type:   TypeValidation,
		Errors: fields,
	}
}

func (e *Error) Error() string {
	if len(e.Errors) > 0 && e.Errors[0] != nil {
		return e.Errors[0].Error
	}
	return "unknown error"
}
