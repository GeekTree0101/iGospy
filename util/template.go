package util

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

// Template is template object fo server
type Template struct {
	Templates *template.Template
}

// Render is render for echo
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
