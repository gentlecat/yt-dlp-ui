package templates

import (
	"html/template"
	"io"
)

// ExecuteTemplates is a custom template executor that uses our template
// structure. Should be used when rendering templates based on "base.html"
// template.
func ExecuteTemplates(wr io.Writer, data interface{}, filenames ...string) error {
	filenames = append(filenames, "frontend/templates/base.html")
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(wr, "base", data)
}
