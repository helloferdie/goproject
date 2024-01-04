package handler

import (
	"spun/pkg/liberror"

	"github.com/helloferdie/golib/v2/liblocale"
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

// Set response
func (resp *Response) Set(code int, data interface{}, err *liberror.Error) {
	resp.Code = code
	resp.Data = data
	resp.Error = err
}

// Succes return default success response
func (resp *Response) Success(data interface{}, err *liberror.Error) {
	resp.Set(200, data, err)
}

// FormatResponse handling echo labstack response in JSON format
func FormatResponse(c echo.Context, resp *Response) error {
	localizer := GetLocalizer(c)
	json := &ResponseJSON{
		Code: resp.Code,
		Data: resp.Data,
	}

	if resp.Error != nil {
		err := resp.Error.Errors[0]
		json.Message = resp.Error.Type
		json.Error = err.Error
		errLocale, _ := liblocale.Translate(localizer, err.Error, err.ErrorVars)
		json.ErrorLocale = errLocale

		// Parse error from service layer to HTTP status code
		if json.Code == 0 {
			if resp.Error.Type == liberror.TypeValidation {
				json.Code = 422
				errList := []interface{}{}
				for k, err := range resp.Error.Errors {
					if k == 0 {
						errList = append(errList, map[string]interface{}{
							"field":        err.Field,
							"error":        err.Error,
							"error_locale": json.ErrorLocale,
						})
						continue
					}

					errLocale, _ := liblocale.Translate(localizer, err.Error, err.ErrorVars)
					errList = append(errList, map[string]interface{}{
						"field":        err.Field,
						"error":        err.Error,
						"error_locale": errLocale,
					})
				}
				json.Data = errList
			} else {
				switch err.Error {
				case liberror.ErrUnauthorized:
					json.Code = 401
				case liberror.ErrForbidden:
					json.Code = 403
				case liberror.ErrNotFound:
					json.Code = 404
				default:
					json.Code = 500
				}
			}
		}
	}

	if resp.Code == 200 {
		json.Success = true
		json.Message = "common.success.default"
	}

	messageLocale, _ := liblocale.Translate(localizer, json.Message, nil)
	json.MessageLocale = messageLocale
	return c.JSON(json.Code, json)
}
