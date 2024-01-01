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
	ErrNotFound     = "common.error.request.not_found"

	// Server
	ErrRepository     = "common.error.server.repository"
	ErrNotImplemented = "common.error.server.not_implemented"
)

func NewError(errorType string, errorMessages ...string) *Error {
	baseErrors := make([]*Base, len(errorMessages))
	for i, msg := range errorMessages {
		baseErrors[i] = &Base{Error: msg}
	}
	return &Error{
		Type:   errorType,
		Errors: baseErrors,
	}
}

func NewServerError(errMsg string) *Error {
	return NewError(TypeServer, errMsg)
}

func NewRequestError(errMsg string) *Error {
	return NewError(TypeRequest, errMsg)
}

func NewErrNotImplemented() *Error {
	return NewServerError(ErrNotImplemented)
}

func NewErrRepository() *Error {
	return NewServerError(ErrRepository)
}

func NewErrNotFound() *Error {
	return NewRequestError(ErrNotFound)
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
