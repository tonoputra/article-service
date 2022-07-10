package main

import (
	db "article-cache/internal/db/mongo"
	"article-cache/internal/helper"
	"article-cache/internal/routes"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var err error

	// Load environment file
	if os.Getenv("LOAD_DOT_ENV") == "true" {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Load environment variable CONTEXT_TIMEOUT
	_, err = strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error at load environment variable CONTEXT_TIMEOUT. error: %v", err)
	}

	api := echo.New()

	// Custom Binder
	api.Binder = &helper.CustomBinder{}

	// Custom Validator
	api.Validator = &helper.CustomValidator{Validator: validator.New()}

	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	// Cors Middleware for API endpoint.
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	db.Mongo.Init()
	helper.Redis.Init()
	routes.Swagger.Init(api)

	routes.RouteApply(api, "/v1")

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		api.ServeHTTP(res, req)

		return
	})

	server.Logger.Fatal(server.Start(":" + os.Getenv("SERVICE_PORT")))
}
