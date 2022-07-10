package routes

import (
	"os"

	"github.com/labstack/echo/v4" // swagger embed files
	echoSwagger "github.com/swaggo/echo-swagger"
	"article-cache/docs"
)

type SwaggerInterface interface {
	Init(e *echo.Echo)
}

type swagger struct {
}

func (*swagger) Init(e *echo.Echo) {
	docs.SwaggerInfo.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo.Description = os.Getenv("SWAGGER_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("SWAGGER_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SWAGGER_SCHEME")}

	api := e.Group("swagger")
	api.GET("/*", echoSwagger.WrapHandler)
}

var Swagger = &swagger{}
