package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func downloadRoot(c echo.Context) (err error) {
	files := []os.FileInfo{}
	err = filepath.Walk("uploads", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || strings.HasPrefix(info.Name(), "_") {
			return nil
		}
		files = append(files, info)
		return nil
	})
	c.Render(http.StatusOK, "fileserver", files)
	return
}

func downloader(c echo.Context) error {
	path := filepath.Join("uploads", c.Param("file"))
	//bucket := ratelimit.NewBucketWithRate(1024*1024, 1024*1024)
	//file, err := os.Open(path)
	//if err != nil {
	//	return err
	//}
	//_, err = io.Copy(c.Response().Writer, ratelimit.Reader(file, bucket))
	return c.Attachment(path, c.Param("file"))
}
