package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	fmt.Printf("error: %s\n", err.Error())
}
func (a *application) routes() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.StaticFS("/static", a.files.Static)
	e.Renderer = NewEmbedTemplateRenderer(a.files.Templates)

	//Setup Routes
	e.GET("/", a.templateHome)
	e.GET("/bills", a.templateBillsList)
	e.POST("/bills", a.templateAddBill)
	e.GET("/bills/new", a.templateBillForm)
	e.PUT("/bills/:id", a.templateUpdateBill)
	e.GET("/bills/:id/edit", a.templateGetEditBill)
	e.DELETE("/bills/:id", a.templateDeleteBill)
	e.GET("/ws/hotreload", a.websocketUpgrade)
	return e
}
