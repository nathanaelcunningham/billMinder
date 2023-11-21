package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func (a *application) websocketUpgrade(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	// Your WebSocket logic here, such as reading and writing messages, etc.
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if err = ws.WriteMessage(websocket.TextMessage, a.uuid); err != nil {
			break
		}
	}
	return nil
}
