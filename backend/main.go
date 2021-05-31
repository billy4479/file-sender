package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.String("p", "4479", "port")
	flag.Parse()

	e := echo.New()

	e.Static("/", "../dist/frontend")
	e.GET("/upload", getID)
	e.GET("/status", ws)
	e.POST("/upload/:id", upload)
	e.GET("/download/:id", download)

	e.Use(middleware.Secure())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${remote_ip}]: ${method} on ${uri} -> Got ${status} in ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	err := e.Start(":" + *port)
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
