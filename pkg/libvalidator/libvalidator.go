package libvalidator

import (
	"reflect"
	"spun/pkg/liberror"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		json := f.Tag.Get("json")
		loc := f.Tag.Get("loc")
		if strings.HasSuffix(loc, ".") {
			loc += json
		}
		return json + "|" + loc
	})
}

// Validate
func Validate(obj interface{}) *liberror.Error {
	err := validate.Struct(obj)
	if err == nil {
		// Validation check passed, no validation error
		return nil
	}

	errFields := []*liberror.Base{}
	errValidators := err.(validator.ValidationErrors)
	for _, errValidator := range errValidators {
		// Split tag name "json|loc" from given struct, check RegisterTagNameFunc
		errType := errValidator.Tag()
		tag := strings.Split(errValidator.Field(), "|")
		jsonTag := tag[0]
		locTag := tag[1]

		// Prepare error locale variables
		errVars := map[string]string{
			"field": "{{" + locTag + "}}",
		}
		switch errType {
		case "min":
		case "max":
			errVars["val"] = errValidator.Param()
		}

		// Append error locale
		errFields = append(errFields, &liberror.Base{
			Error:     "common.error.validation." + errType,
			Field:     jsonTag,
			ErrorVars: errVars,
		})
	}
	return liberror.NewErrValidation(errFields...)
}
