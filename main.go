package main

import (
	"context"
	"flag"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.String("p", "4479", "port")
	flag.Parse()
	err := os.MkdirAll("uploads", 0755)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	e.Static("/", "public")
	e.POST("/upload", upload)
	e.GET("/download", downloadRoot)
	e.GET("/download/:file", downloader)

	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	err = e.Start(":" + *port)
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func upload(c echo.Context) (err error) {
	h, err := c.FormFile("file")
	if err != nil {
		return
	}

	file, err := os.Create(filepath.Join("uploads", h.Filename))
	if err != nil {
		return
	}

	f, err := h.Open()
	if err != nil {
		return
	}

	//bucket := ratelimit.NewBucketWithRate(1024*1024, 1024*1024)

	//_, err = io.Copy(file, ratelimit.Reader(f, bucket))
	_, err = io.Copy(file, f)
	if err != nil {
		return
	}
	file.Close()

	return c.Redirect(http.StatusCreated, "/")
}
