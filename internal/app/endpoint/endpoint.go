package endpoint

import (
	"github.com/labstack/echo/v4"
	"github.com/yanasirina/umbrella_corp_test_task/internal/app/service"
	"net/http"
)

type Endpoint struct {
	service service.Service
}

func New() *Endpoint {
	return &Endpoint{}
}

type PingResponse struct {
	DaysBefore2025 int
}

func (endpoint *Endpoint) PingEndpoint(context echo.Context) error {
	daysBefore2025 := endpoint.service.GetDaysBefore2025()
	response := PingResponse{DaysBefore2025: daysBefore2025}
	return context.JSON(http.StatusOK, response)
}
