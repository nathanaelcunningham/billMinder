package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nathanaelcunningham/billReminder/components"
	"github.com/nathanaelcunningham/billReminder/models"
)

func (a *application) home(c echo.Context) error {
	comp := components.Layout()
	buf := bytes.NewBuffer(nil)
	comp.Render(c.Request().Context(), buf)

	return c.HTML(200, buf.String())
}

func (a *application) billsTable(c echo.Context) error {
	bills, err := a.billRepo.GetAll()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return err
	}
	total := 0.0
	for _, bill := range bills {
		total += bill.Amount
	}
	comp := components.BillTable(bills, total)
	buf := bytes.NewBuffer(nil)
	comp.Render(c.Request().Context(), buf)

	return c.HTML(200, buf.String())
}

func (a *application) addBill(c echo.Context) error {
	billCreate := models.Bill{}
	billCreate.Name = c.FormValue("name")
	amt, _ := strconv.ParseFloat(c.FormValue("amount"), 64)
	billCreate.Amount = amt
	day, _ := strconv.ParseInt(c.FormValue("dueDateDay"), 10, 64)
	billCreate.DueDateDay = day

	id, err := a.billRepo.Create(&billCreate)
	if err != nil {
		return c.String(500, err.Error())
	}

	bill, err := a.billRepo.Get(id)

	comp := components.AddBill(*bill)
	buf := bytes.NewBuffer(nil)
	comp.Render(c.Request().Context(), buf)

	return c.HTML(200, buf.String())
}

func (a *application) getBillRow(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	bill, err := a.billRepo.Get(id)
	if err != nil {
		return c.String(500, err.Error())
	}

	comp := components.GetBillRow(*bill)
	buf := bytes.NewBuffer(nil)
	comp.Render(c.Request().Context(), buf)

	return c.HTML(200, buf.String())
}
func (a *application) updateBillRow(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	billCreate := models.Bill{
		ID: id,
	}
	billCreate.Name = c.FormValue("name")
	amt, _ := strconv.ParseFloat(c.FormValue("amount"), 64)
	billCreate.Amount = amt
	day, _ := strconv.ParseInt(c.FormValue("dueDateDay"), 10, 64)
	billCreate.DueDateDay = day

	err := a.billRepo.Update(&billCreate)
	if err != nil {
		return c.String(500, err.Error())
	}

	bill, err := a.billRepo.Get(id)
	if err != nil {
		return c.String(500, err.Error())
	}

	comp := components.GetBillRow(*bill)
	buf := bytes.NewBuffer(nil)
	comp.Render(c.Request().Context(), buf)

	return c.HTML(200, buf.String())
}
func (a *application) editBillRow(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	bill, err := a.billRepo.Get(id)
	if err != nil {
		return c.String(500, err.Error())
	}

	comp := components.EditBillRow(*bill)
	buf := bytes.NewBuffer(nil)
	comp.Render(c.Request().Context(), buf)

	return c.HTML(200, buf.String())
}
