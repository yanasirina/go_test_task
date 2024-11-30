package endpoint

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	GetDaysBefore2025() int
}

type Endpoint struct {
	service Service
}

func New(service Service) *Endpoint {
	return &Endpoint{
		service: service,
	}
}

type PingResponse struct {
	DaysBefore2025 int
}

func (endpoint *Endpoint) PingEndpoint(context echo.Context) error {
	daysBefore2025 := endpoint.service.GetDaysBefore2025()
	response := PingResponse{DaysBefore2025: daysBefore2025}
	return context.JSON(http.StatusOK, response)
}
