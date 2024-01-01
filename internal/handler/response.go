package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spun/pkg/liberror"
	"spun/pkg/liberror/mapper"

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

// // NewResponse
// func NewResponse() *Response {
// 	return &Response{
// 		Code:  501,
// 		Data:  nil,
// 		Error: liberror.NewErrNotImplemented(),
// 	}
// }

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

// // ErrorResponse
// func (resp *Response) ErrorResponse(data interface{}, err *liberror.Error) {
// 	resp
// }

// ErrorHandler handling echo labstack default error
func ErrorHandler(err error, c echo.Context) {
	//code := http.StatusInternalServerError
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := new(Response)
	resp.Code = he.Code
	resp.Error = mapper.MapHTTPError(he.Code)

	if he.Code == 400 {
		if e1, ok := he.Internal.(*json.UnmarshalTypeError); ok {
			resp.Error.Errors[0].ErrorVars = map[string]interface{}{
				"f": e1.Field,
				"e": e1.Type,
				"v": e1.Value,
			}
		}
	}

	// switch code {
	// case http.StatusNotFound:
	// 	resp.Error = liberror.NewRequestError(liberror.ErrRouteNotFound)
	// case http.StatusUnsupportedMediaType:
	// 	resp.Error = liberror.NewRequestError(liberror.ErrMediaType)
	// }

	fmt.Println(err)

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
