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
	Teeth int
}

type Contact struct {
	Email string
	Name  string
}

func newContact(name string, email string) Contact {
	return Contact{
		Email: name,
		Name:  email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func (d Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Value  map[string][]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Value:  make(map[string][]string),
		Errors: make(map[string]string),
	}
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

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func main() {
	e := echo.New()
	data := newData()
	e.Use(middleware.Logger())

	page := newPage()
	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {

		return c.Render(200, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Value["name"] = []string{name}
			formData.Value["email"] = []string{email}
			formData.Errors["email"] = "Email already exists"
			return c.Render(400, "form", data)

		}
		page.Data.Contacts = append(page.Data.Contacts, newContact(name, email))
		return c.Render(200, "display", page)
	})

	e.Logger.Fatal(e.Start(":42049"))

}
