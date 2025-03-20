package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Route struct {
}

func RegisterRouters(e *echo.Echo) {
	e.Use(middleware.CORS())

	InitCharacterRoute(e)
}
