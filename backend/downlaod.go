package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func download(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("Invalid ID")
	}

	callback, ok := callbacks[id]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if callback.called {
		return echo.NewHTTPError(http.StatusBadRequest, "This file was already downloaded")
	}

	callback.called = true
	err := callback.dlCallback(c.Response().Writer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
