package response

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type (
	Body struct {
		Status         string      `json:"status"`
		HTTPStatusCode int         `json:"-"`
		Message        string      `json:"message,omitempty"`
		Total          int         `json:"total,omitempty" swaggerignore:"true"`
		SubTotal       int         `json:"subtotal,omitempty" swaggerignore:"true"`
		Meta           interface{} `json:"meta,omitempty" swaggerignore:"true"`
		Data           interface{} `json:"values,omitempty" swaggerignore:"true"`
		Errors         interface{} `json:"errors,omitempty" swaggerignore:"true"`
	}
)

// WriteSuccess as JSON
func WriteSuccess(ctx echo.Context, b Body) error {
	if b.Status == "" {
		b.Status = "success"
	}

	return ctx.JSON(http.StatusOK, b)
}

// WriteError as JSON
func WriteError(ctx echo.Context, b Body) error {
	if b.Status == "" {
		b.Status = "error"
	}

	// Error mapping messages
	if b.Errors != nil {

		switch v := b.Errors.(type) {
		case *echo.HTTPError:
			b.Errors = v
		case *url.Error:
			e := map[string]interface{}{
				"method":  v.Op,
				"url":     v.URL,
				"message": v.Err.Error(),
			}
			b.Errors = e
		case error:
			e := fmt.Sprintf("%v", v)
			b.Errors = map[string]interface{}{
				"message": e,
			}
		default:
			ctx.Logger().Errorf("WriteError stackEntries type is: %v", v)
		}

	}

	return ctx.JSON(b.HTTPStatusCode, b)
}

func WriteErrorBinding(ctx echo.Context, err error) {
	WriteError(ctx, Body{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        "Error binding payload data",
		Errors:         err,
	})
}

func WriteErrorValidate(ctx echo.Context, err error) {
	WriteError(ctx, Body{
		HTTPStatusCode: http.StatusBadRequest,
		Message:        "Error validate payload data",
		Errors:         err,
	})
}
