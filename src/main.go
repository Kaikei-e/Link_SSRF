package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	port := ":9000"
	t := &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return getHomePage(c)
	})

	// e.POST("/", func(c echo.Context) error {
	// 	return getHomePage(c)
	// })

	e.Logger.Fatal(e.Start(port))
}

func getHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "hoge")
}
