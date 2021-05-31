package main

import (
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func upload(c echo.Context) error {
	h, err := c.FormFile("file")
	if err != nil {
		return err
	}

	id := c.Param("id")

	wg := sync.WaitGroup{}
	wg.Add(1)
	callbackData := makeCallback(id, &wg, h)

	wg.Wait()

	if strings.HasPrefix(callbackData.statusCallback(), "Error: ") {
		return echo.NewHTTPError(http.StatusInternalServerError, callbackData.statusCallback())
	}
	return c.String(http.StatusOK, callbackData.statusCallback())
}

func getID(c echo.Context) error {
	id := uuid.NewString()
	for _, ok := callbacks[id]; ok; {
		id = uuid.NewString()
	}

	return c.String(http.StatusOK, id)
}
