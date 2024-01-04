package handler

import (
	"encoding/json"
	"net/http"
	"spun/pkg/liberror"

	"github.com/labstack/echo/v4"
)

// ErrorHandler handling echo labstack default error
func ErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	errMsg := he.Message.(string)

	resp := new(Response)
	resp.Code = he.Code

	switch he.Code {
	case 400:
		if e1, ok := he.Internal.(*json.UnmarshalTypeError); ok {
			resp.Error = liberror.NewError("common.error.request.bad", []*liberror.Base{{
				Error: "common.error.request.unmarshal",
				ErrorVars: map[string]interface{}{
					"f": e1.Field,
					"e": e1.Type,
					"v": e1.Value,
				},
			}})
		} else {
			resp.Error = liberror.NewErrorQuick("common.error.request.bad", errMsg)
		}
	case 401:
		resp.Error = liberror.NewErrorQuick("common.error.request.unauthorized", errMsg)
	case 403:
		resp.Error = liberror.NewErrorQuick("common.error.request.forbidden", errMsg)
	case 404:
		resp.Error = liberror.NewErrorQuick("common.error.request.not_found.default", "common.error.request.not_found.route")
	case 405:
		resp.Error = liberror.NewErrorQuick("common.error.request.method", "common.error.request.method")
	case 408:
		resp.Error = liberror.NewErrorQuick("common.error.request.timeout", "common.error.request.timeout")
	case 415:
		resp.Error = liberror.NewErrorQuick("common.error.request.media_type", "common.error.request.media_type")
	case 429:
		resp.Error = liberror.NewErrorQuick("common.error.request.many", "common.error.request.many")
	case 500:
		resp.Error = liberror.NewErrorQuick("common.error.server.default", "common.error.server.default")
	case 501:
		resp.Error = liberror.NewErrorQuick("common.error.server.not_implemented", "common.error.server.not_implemented")
	case 504:
		resp.Error = liberror.NewErrorQuick("common.error.server.timeout", "common.error.server.timeout")
	default:
		resp.Error = liberror.NewErrorQuick("common.error.server.unknown", errMsg)
	}
	FormatResponse(c, resp)
}
