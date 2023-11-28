package gomponents

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"github.com/nathanaelcunningham/billReminder/models"
)

func BillList(bills []models.Bill) g.Node {
	return Div(
		ID("bill-list"),
		Class("flex flex-col gap-y-4 pt-4"),
		g.Group(g.Map(bills, func(bill models.Bill) g.Node {
			return BillCard(bill)
		})),
	)
}

func BillCard(bill models.Bill) g.Node {
	return Div(
		Div(
			ID("bill-row"),
			Class("relative bg-white rounded-sm shadow-md shadow-slate-200"),
			hx.Get(fmt.Sprintf("/bills/%d/edit", bill.ID)),
			hx.Target("#bill-list"),
			hx.Swap("innerHTML"),
			Div(Class("flex justify-between py-6 px-4"),
				Div(
					Class("flex flex-col gap-y-1"),
					P(Class("text-xl"), g.Text(bill.Name)),
					P(Class("text-sm text-slate-500"), g.Text(getDate(bill.DueDateDay))),
				),
				Div(
					Class("flex flex-col justify-center"),
					Div(g.Text(moneyFmt(bill.Amount))),
				),
			),
			Div(
				Class("absolute top-0 right-0 mt-2 mr-4"),
				Div(Class("flex gap-x-2 items-center"),
					Div(
						c.Classes{
							"px-1 text-xs text-white lowercase rounded-lg": true,
							"bg-slate-400": bill.BillType == models.STATIC,
							"bg-blue-400":  bill.BillType == models.VARIABLE,
						},
						g.Text(string(bill.BillType)),
					),
					g.If(bill.IsAutoPay, I(Class("gg-bolt text-sky-700"))),
				),
			),
		),
	)
}

func BillForm(bill models.Bill) g.Node {
	return Div(
		ID("bill-form"),
		Class("grid grid-cols-1 gap-2"),
		Input(
			Class("p-2 rounded border border-zinc-200"),
			Type("text"),
			Name("name"),
			Placeholder("Bill Name"),
			Value(bill.Name),
		),
		Div(
			Class("grid grid-cols-6 gap-2"),
			Select(
				Class("col-span-4 p-2 rounded border border-zinc-200"),
				Name("dueDateDay"),
				g.Group(g.Map(numList(1, 31), func(num int) g.Node {
					return Option(
						g.If(num == int(bill.DueDateDay), g.Attr("selected", "selected")),
						Value(fmt.Sprintf("%d", num)),
						g.Text(fmt.Sprintf("%d", num)),
					)
				})),
			),
			Div(
				Class("flex col-span-2 gap-2 justify-self-end items-center"),
				Label(For("isAutoPay"), g.Text("Auto Pay")),
				Input(
					ID("isAutoPay"),
					Type("checkbox"),
					Name("isAutoPay"),
					g.If(bill.IsAutoPay, g.Attr("checked", "checked")),
				),
			),
		),
		Input(
			Class("p-2 rounded border border-zinc-200"),
			Type("number"),
			Step("0.01"),
			Name("amount"),
			Placeholder("Bill Amount"),
			Value(fmt.Sprintf("%.2f", bill.Amount)),
		),
		Select(
			Class("p-2 rounded border border-zinc-200"),
			Name("billType"),
			g.Group(
				g.Map(
					[]models.BillType{models.STATIC, models.VARIABLE},
					func(billType models.BillType) g.Node {
						return Option(
							g.If(billType == bill.BillType, g.Attr("selected", "selected")),
							Value(string(billType)),
							g.Text(string(billType)),
						)
					},
				),
			),
		),
		Div(
			Class("flex flex-row gap-x-4 justify-between"),
			Div(
				g.If(bill.ID != 0,
					Button(
						Class(
							"py-2 px-4 rounded border border-red-600 hover:text-white hover:bg-red-600",
						),
						hx.Delete(fmt.Sprintf("/bills/%d", bill.ID)),
						hx.Target("#bill-list"),
						hx.Swap("outerHTML"),
						HScript("on htmx:afterRequest remove #bill-form"),
						g.Text("Delete"),
					),
				),
			),
			Div(
				Class("flex gap-x-1 justify-end"),
				Button(
					Class(
						"py-2 px-4 rounded border hover:text-white border-zinc-400 hover:bg-zinc-400",
					),
					hx.Get("/bills"),
					hx.Target("#bill-list"),
					hx.Swap("outerHTML"),
					g.Text("Cancel"),
				),
				Button(
					Class(
						"py-2 px-4 rounded border hover:text-white border-sky-600 text-slate-600 hover:bg-sky-600",
					),
					g.If(bill.ID == 0, hx.Post("/bills")),
					hx.Put(fmt.Sprintf("/bills/%d", bill.ID)),
					hx.Include("closest #bill-form"),
					hx.Target("#bill-list"),
					hx.Swap("outerHTML"),
					HScript("on click wait 500ms then remove closest #bill-form"),
					g.Text("Save"),
				),
			),
		),
	)
}
