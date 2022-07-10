package helper

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// use a single instance , it caches struct info
var uni *ut.UniversalTranslator

func (cv *CustomValidator) Validate(i interface{}) error {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	cv.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})

	en_translations.RegisterDefaultTranslations(cv.Validator, trans)

	if err := cv.Validator.Struct(i); err != nil {
		var errors ValidationErrors

		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   strings.ToLower(e.Field()),
				Message: strings.ToLower(strings.Replace(e.Translate(trans), "_", " ", 1)),
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, errors)
	}

	return nil
}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}

	return
}
