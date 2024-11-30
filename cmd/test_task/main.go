package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/endpoint"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/middleware"
	service2 "github.com/yanasirina/umbrella_corp_test_task/internal/app/service"
)

func main() {
	server := echo.New()
	server.Use(middleware.New().RedButtonMiddleware)

	service := service2.New()
	server.GET("/ping", endpoint.New(service).PingEndpoint)

	server.Logger.Fatal(server.Start(":1323"))
}
