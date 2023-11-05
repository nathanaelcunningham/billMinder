package main

import (
	"github.com/labstack/echo/v4"
)

func (a *application) routes() *echo.Echo {
	e := echo.New()
	e.GET("/", a.home)
	return e
}
