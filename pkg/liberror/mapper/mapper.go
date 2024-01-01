package mapper

import (
	"net/http"
	"spun/pkg/liberror"
)

// MapHTTPError
func MapHTTPError(code int) *liberror.Error {
	switch code {
	case http.StatusBadRequest: // 400
		return liberror.NewRequestError("common.error.request.bad")
	case http.StatusUnauthorized: // 401
		return liberror.NewRequestError("common.error.request.unauthorized")
	case http.StatusForbidden: // 403
		return liberror.NewRequestError("common.error.request.forbidden")
	case http.StatusNotFound: // 404
		return liberror.NewRequestError("common.error.request.route.not_found")
	case http.StatusMethodNotAllowed: // 405
		return liberror.NewRequestError("common.error.request.method")
	case http.StatusRequestTimeout: // 408
		return liberror.NewRequestError("common.error.request.timeout")
	case http.StatusInternalServerError: // 500
		return liberror.NewServerError("common.error.server.default")
	case http.StatusNotImplemented: // 501
		return liberror.NewServerError("common.error.server.default")
	case http.StatusGatewayTimeout: // 504
		return liberror.NewServerError("common.error.server.timeout")
	default:
		errType := liberror.TypeServer
		if code >= 400 && code < 500 {
			errType = liberror.TypeRequest
		}

		return &liberror.Error{
			Type: errType,
			Errors: []*liberror.Base{
				{Error: http.StatusText(code)},
			},
		}
	}
}
