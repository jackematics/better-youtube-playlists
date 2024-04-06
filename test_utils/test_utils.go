package test_utils

import (
	"bytes"
	"html/template"

	"github.com/jackematics/better-youtube-playlists/helper/func_map"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func ResetServerState() {
	page_data.IndexState = page_data.InitialiseState()
}

var test_func_map = template.FuncMap{
	"getItemWithNumber": func(playlistItem model.PlaylistItem, arrayIndex int) func_map.ItemWithNumber {
		return func_map.ItemWithNumber{
			Id:         playlistItem.Id,
			Title:      playlistItem.Title,
			Thumbnail:  playlistItem.Thumbnail,
			ItemNumber: arrayIndex + 1,
		}
	},
}

func ParseTemplateToString(templateName string, paths []string, state any) string {
	tmpl := template.Must(template.New("test").Funcs(test_func_map).ParseFiles(paths...))

	var result bytes.Buffer
	tmpl.ExecuteTemplate(&result, templateName, state)

	return result.String()
}
