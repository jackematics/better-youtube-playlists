package test_utils

import (
	"bytes"
	"html/template"

	"github.com/jackematics/better-youtube-playlists/lib/func_map"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func ResetServerState() {
	page_data.IndexState = page_data.InitialiseState()
}

func ParseTemplateToString(templateName string, paths []string, state any) string {
	tmpl := template.Must(template.New("test").Funcs(func_map.PageFuncs).ParseFiles(paths...))

	var result bytes.Buffer
	tmpl.ExecuteTemplate(&result, templateName, state)

	return result.String()
}
