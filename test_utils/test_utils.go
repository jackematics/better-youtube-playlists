package test_utils

import (
	"bytes"
	"html/template"
)

func ParseTemplateToString(templateName string, paths []string, state any) string {
	tmpl := template.Must(template.New("test").ParseFiles(paths...))

	var result bytes.Buffer
	tmpl.ExecuteTemplate(&result, templateName, state)

	return result.String()
}
