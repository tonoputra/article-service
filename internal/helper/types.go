package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	Route struct {
		Method      string
		Pattern     string
		HandlerFunc echo.HandlerFunc
		Middleware  []echo.MiddlewareFunc
	}

	CustomBinder struct{}

	CustomValidator struct {
		Validator *validator.Validate
	}

	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationErrors []ValidationError
)
