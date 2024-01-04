package handler

import (
	"encoding/json"
	"net/http"
	"spun/pkg/liberror"

	"github.com/labstack/echo/v4"
)

// ErrorHandler custom error handler for the Echo framework.
// Handle any error thrown during HTTP request processing and formats it into a standardized JSON response.
func ErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	errMsg := he.Message.(string)

	resp := new(Response)
	resp.Code = he.Code

	switch he.Code {
	case 400: // Handle bad request errors.
		if e1, ok := he.Internal.(*json.UnmarshalTypeError); ok {
			resp.Error = liberror.NewError("common.error.request.bad", []*liberror.Base{{
				Error: "common.error.request.unmarshal",
				ErrorVars: map[string]string{
					"f": e1.Field,
					"e": e1.Type.String(),
					"v": e1.Value,
				},
			}})
		} else {
			resp.Error = liberror.NewErrorQuick("common.error.request.bad", errMsg)
		}
	case 401: // Handle unauthorized errors.
		resp.Error = liberror.NewErrorQuick("common.error.request.unauthorized", errMsg)
	case 403: // Handle forbidden errors.
		resp.Error = liberror.NewErrorQuick("common.error.request.forbidden", errMsg)
	case 404: // Handle not found errors.
		resp.Error = liberror.NewErrorQuick("common.error.request.not_found.default", "common.error.request.not_found.route")
	case 405: // Handle method not allowed errors.
		resp.Error = liberror.NewErrorQuick("common.error.request.method", "common.error.request.method")
	case 408: // Handle request timeout errors.
		resp.Error = liberror.NewErrorQuick("common.error.request.timeout", "common.error.request.timeout")
	case 415: // Handle unsupported media type errors.
		resp.Error = liberror.NewErrorQuick("common.error.request.media_type", "common.error.request.media_type")
	case 429: // Handle too many requests errors (rate limiting).
		resp.Error = liberror.NewErrorQuick("common.error.request.many", "common.error.request.many")
	case 500: // Handle internal server errors.
		resp.Error = liberror.NewErrorQuick("common.error.server.default", "common.error.server.default")
	case 501: // Handle not implemented errors.
		resp.Error = liberror.NewErrorQuick("common.error.server.not_implemented", "common.error.server.not_implemented")
	case 504: // Handle gateway timeout errors.
		resp.Error = liberror.NewErrorQuick("common.error.server.timeout", "common.error.server.timeout")
	default: // Handle all other errors.
		resp.Error = liberror.NewErrorQuick("common.error.server.unknown", errMsg)
	}
	FormatResponse(c, resp)
}
