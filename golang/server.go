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

type Contact struct {
	Email   string
	Name string
}

func newContact(name string, email string) Contact {
	return Contact{
		Email: name,
		Name: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@gmail.com"),
			newContact("Jane", "jc@gmail.com"),
		},
	
	}
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
	data := newData()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {

		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		data.Contacts = append(data.Contacts, newContact(name, email))
		return c.Render(200, "index", data)
	})

	e.Logger.Fatal(e.Start(":42049"))

}
