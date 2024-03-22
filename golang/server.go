package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}
type Count struct {
	Count int
	Teeth  int
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("../public/*.html")),
	}
}

func main() {
	e := echo.New()
	count := Count{Count: 0, Teeth: 0}
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {

		return c.Render(200, "index", count)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		count.Teeth += 2
		return c.Render(200, "index", count)
	})

	e.Logger.Fatal(e.Start(":42049"))

}
