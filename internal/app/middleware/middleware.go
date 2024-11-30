package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

type Middleware struct {
}

func New() *Middleware {
	return &Middleware{}
}

const roleAdmin = "admin"

func (middleware *Middleware) RedButtonMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		headers := context.Request().Header
		userRole := headers.Get("User-Role")
		if userRole == roleAdmin {
			log.Println("red button user detected")
		}
		return next(context)
	}
}
