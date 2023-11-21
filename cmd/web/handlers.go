package main

import "github.com/labstack/echo/v4"

func (a *application) dbDownload(c echo.Context) error {
	pass := c.FormValue("password")
	if pass != "abc123!" {
		return c.String(401, "Unauthorized")
	}
	return c.Attachment(a.cfg.DBPath, "billReminder.db")
}
