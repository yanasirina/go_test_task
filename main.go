package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type PingResponse struct {
	DaysBefore2025 int
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func GetDaysBefore2025() int {
	neededDate := Date(2025, 1, 1)
	durationToDate := time.Until(neededDate)
	daysBeforeDate := int(durationToDate.Hours()) / 24
	return daysBeforeDate
}

func PingHandler(context echo.Context) error {
	daysBefore2025 := GetDaysBefore2025()
	response := PingResponse{DaysBefore2025: daysBefore2025}
	return context.JSON(http.StatusOK, response)
}

func RedButtonMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		headers := context.Request().Header
		userRole := headers.Get("User-Role")
		if userRole == "admin" {
			log.Println("red button user detected")
		}
		return next(context)
	}
}

func main() {
	server := echo.New()
	server.Use(RedButtonMiddleware)
	server.GET("/ping", PingHandler)
	server.Logger.Fatal(server.Start(":1323"))
}
