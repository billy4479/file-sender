package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ws(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	data, ok := callbacks[id]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	currentStatus := data.statusCallback()
	_, err = writeMessage(ws, currentStatus)
	if err != nil {
		fmt.Println(err)
	}

	for {
		if currentStatus != data.statusCallback() {
			currentStatus = data.statusCallback()
			shouldStop, err := writeMessage(ws, currentStatus)
			if err != nil {
				fmt.Println(err)
			}
			if shouldStop {
				break
			}
		}

		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func writeMessage(ws *websocket.Conn, message string) (shouldStop bool, err error) {
	shouldStop = false
	if message == "Done" || strings.HasPrefix(message, "Error: ") {
		shouldStop = true
		err = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, message))
		return
	}
	err = ws.WriteMessage(websocket.TextMessage, []byte(message))
	return
}
