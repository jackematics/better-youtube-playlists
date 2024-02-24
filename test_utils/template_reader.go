package test_utils

import (
	"bytes"
	"html/template"
	"os"

	"github.com/jackematics/better-youtube-playlists/model"
)

func ParseTemplateToString(path string) string {
	htmlBytes, _ := os.ReadFile(path)
	htmlString := string(htmlBytes)
	tmpl, _ := template.New("html").Parse(htmlString)
	var expectedHtml bytes.Buffer
	tmpl.Execute(&expectedHtml, model.ModalModel{Hidden: false})

	return expectedHtml.String()
}
