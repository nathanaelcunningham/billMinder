package main

import (
	"github.com/labstack/echo/v4"
)

func (a *application) routes() *echo.Echo {
	e := echo.New()
	e.GET("/", a.home)
	e.GET("/bills", a.billsTable)
	e.POST("/bills", a.addBill)
	e.GET("/bills/:id", a.getBillRow)
	e.PUT("/bills/:id", a.updateBillRow)
	e.GET("/bills/:id/edit", a.editBillRow)
	return e
}
