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

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		daysBefore2025 := GetDaysBefore2025()
		response := PingResponse{DaysBefore2025: daysBefore2025}
		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
