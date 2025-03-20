package main

import (
	v1 "github.com/ionnotion/statictools-server/api/v1"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	v1.RegisterRouters(e)

}
