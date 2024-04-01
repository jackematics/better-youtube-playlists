package test_utils

import (
	"bytes"
	"html/template"

	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func ResetServerState() {
	page_data.IndexState = page_data.InitialiseState()
}

func ParseTemplateToString(templateName string, paths []string, state any) string {
	tmpl := template.Must(template.ParseFiles(paths...))

	var result bytes.Buffer
	tmpl.ExecuteTemplate(&result, templateName, state)

	return result.String()
}
