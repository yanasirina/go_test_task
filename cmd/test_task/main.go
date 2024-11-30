package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/endpoint"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/middleware"
)

func main() {
	server := echo.New()
	server.Use(middleware.New().RedButtonMiddleware)
	server.GET("/ping", endpoint.New().PingEndpoint)
	server.Logger.Fatal(server.Start(":1323"))
}
