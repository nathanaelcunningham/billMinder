package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nathanaelcunningham/billReminder/models"
)

func (a *application) templateHome(c echo.Context) error {

	return c.Render(http.StatusOK, "base.html", nil)
}
func (a *application) templateBillForm(c echo.Context) error {
	return c.Render(http.StatusOK, "bills/billForm.html", models.Bill{})
}

func (a *application) templateBillsList(c echo.Context) error {
	bills, err := a.billRepo.GetAll()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return err
	}
	total := 0.0
	for _, bill := range bills {
		total += bill.Amount
	}
	data := struct {
		Bills []models.Bill
		Total float64
	}{
		Bills: bills,
		Total: math.Round(total*100) / 100,
	}

	return c.Render(http.StatusOK, "bills/billList.html", data)
}

func (a *application) templateGetEditBill(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	bill, err := a.billRepo.Get(id)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.Render(http.StatusOK, "bills/billForm.html", bill)
}

func (a *application) templateAddBill(c echo.Context) error {
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
	billCreate.BillType = models.BillType(c.FormValue("billType"))

	_, err := a.billRepo.Create(&billCreate)
	if err != nil {
		return c.String(500, err.Error())
	}
	return a.templateBillsList(c)
}

func (a *application) templateUpdateBill(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	billUpdate := models.Bill{
		ID: id,
	}
	billUpdate.Name = c.FormValue("name")
	amt, _ := strconv.ParseFloat(c.FormValue("amount"), 64)
	billUpdate.Amount = amt
	day, _ := strconv.ParseInt(c.FormValue("dueDateDay"), 10, 64)
	billUpdate.DueDateDay = day
	billUpdate.IsAutoPay = c.FormValue("isAutoPay") == "on"
	billUpdate.BillType = models.BillType(c.FormValue("billType"))

	err := a.billRepo.Update(&billUpdate)
	if err != nil {
		return c.String(500, err.Error())
	}

	return a.templateBillsList(c)
}
func (a *application) templateDeleteBill(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := a.billRepo.Delete(id)
	if err != nil {
		return c.String(500, err.Error())
	}
	return a.templateBillsList(c)
}
