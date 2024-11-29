package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingResponse struct {
	DaysBefore2025 int
}

func GetDaysBefore2025() int {
	return 5
}

func PingHandler(context echo.Context) error {
	daysBefore2025 := GetDaysBefore2025()
	response := PingResponse{DaysBefore2025: daysBefore2025}
	return context.JSON(http.StatusOK, response)
}

func main() {
	server := echo.New()
	server.GET("/ping", PingHandler)
	server.Logger.Fatal(server.Start(":1323"))
}
