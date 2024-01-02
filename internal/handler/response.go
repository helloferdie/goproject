package handler

import (
	"encoding/json"
	"net/http"
	"spun/pkg/liberror"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code  int
	Data  interface{}
	Error *liberror.Error
}

type ResponseJSON struct {
	Success       bool        `json:"success"`
	Code          int         `json:"code"`
	Message       string      `json:"message"`
	MessageLocale string      `json:"message_locale"`
	Error         string      `json:"error"`
	ErrorLocale   string      `json:"error_locale"`
	Data          interface{} `json:"data"`
}

// Set
func (resp *Response) Set(code int, data interface{}, err *liberror.Error) {
	resp.Code = code
	resp.Data = data
	resp.Error = err
}

// Succes
func (resp *Response) Success(data interface{}, err *liberror.Error) {
	resp.Set(200, data, err)
}

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

// FormatResponse -
func FormatResponse(c echo.Context, resp *Response) error {
	json := &ResponseJSON{
		Code: resp.Code,
		Data: resp.Data,
	}

	if resp.Error != nil {
		err := resp.Error.Errors[0]
		json.Message = resp.Error.Type
		json.Error = err.Error

		// Parse error to HTTP code
		if json.Code == 0 {
			if resp.Error.Type == liberror.TypeValidation {
				json.Code = 422
				json.Data = resp.Error.Errors
			} else {
				switch err.Error {
				case liberror.ErrNotFound:
					json.Code = 404
				}
			}
		}
	}

	if resp.Code == 200 {
		json.Success = true
		json.Message = "common.success.default"
	}
	return c.JSON(json.Code, json)
}
