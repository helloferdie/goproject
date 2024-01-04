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

func Validate(obj interface{}) *liberror.Error {
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	errFields := []*liberror.Base{}
	errValidators := err.(validator.ValidationErrors)
	for _, errValidator := range errValidators {
		errType := errValidator.Tag()
		tag := strings.Split(errValidator.Field(), "|")
		jsonTag := tag[0]
		locTag := tag[1]

		errFields = append(errFields, &liberror.Base{
			Error: "common.error.validation." + errType,
			Field: jsonTag,
			ErrorVars: map[string]string{
				"field": "{{" + locTag + "}}",
			},
		})
	}
	return liberror.NewErrValidation(errFields...)
}
