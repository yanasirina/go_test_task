package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/endpoint"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/middleware"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/service"
	"log"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	server   *echo.Echo
}

func New() (*App, error) {
	app := &App{}
	app.service = service.New()
	app.endpoint = endpoint.New(app.service)

	app.server = echo.New()
	app.server.Use(middleware.New().RedButtonMiddleware)
	app.server.GET("/ping", app.endpoint.PingEndpoint)

	return app, nil
}

func (app *App) Run() error {
	log.Println("Server running...")
	err := app.server.Start(":1323")
	if err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}
	return nil
}
