package test_utils

import (
	"bytes"
	"html/template"
	"os"

	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func ResetServerState() {
	page_data.IndexState = page_data.InitialiseState()
}

func ParseTemplateToString(path string, state any) string {
	htmlBytes, _ := os.ReadFile(path)
	htmlString := string(htmlBytes)
	tmpl, _ := template.New("html").Parse(htmlString)
	var expectedHtml bytes.Buffer
	tmpl.Execute(&expectedHtml, state)

	return expectedHtml.String()
}
