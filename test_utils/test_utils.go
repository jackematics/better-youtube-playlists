package test_utils

import (
	"bytes"
	"html/template"
	"os"

	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
)

func ResetServerState() {
	page_data_repository.IndexState = page_data_repository.SetInitialState()
}

func ParseTemplateToString(path string, model model.ModalModel) string {
	htmlBytes, _ := os.ReadFile(path)
	htmlString := string(htmlBytes)
	tmpl, _ := template.New("html").Parse(htmlString)
	var expectedHtml bytes.Buffer
	tmpl.Execute(&expectedHtml, model)

	return expectedHtml.String()
}
