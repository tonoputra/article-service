package routes

import (
	"article-cache/internal/helper"
	"article-cache/internal/modules/v1/articles"

	"github.com/labstack/echo/v4"
)

type ApiRoute interface {
	Routes() []helper.Route
}

func RouteApply(e *echo.Echo, pathPrefix string) {
	var routers []helper.Route

	apiRoute := []ApiRoute{
		articles.NewRoute(),
	}

	for _, ar := range apiRoute {
		routers = append(routers, ar.Routes()...)
	}

	api := e.Group("api" + pathPrefix)
	for _, router := range routers {

		switch router.Method {
		case echo.GET:
			{
				api.GET(router.Pattern, router.HandlerFunc, router.Middleware...)
				break
			}
		case echo.POST:
			{
				api.POST(router.Pattern, router.HandlerFunc, router.Middleware...)
				break
			}
		case echo.PUT:
			{
				api.PUT(router.Pattern, router.HandlerFunc, router.Middleware...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(router.Pattern, router.HandlerFunc, router.Middleware...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(router.Pattern, router.HandlerFunc, router.Middleware...)
				break
			}
		}

	}
}
