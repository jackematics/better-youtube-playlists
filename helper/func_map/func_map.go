package func_map

import (
	"text/template"

	"github.com/jackematics/better-youtube-playlists/model"
)

type ItemWithNumber struct {
	Id         string
	Title      string
	Thumbnail  model.Thumbnail
	ItemNumber int
}

var getItemWithNumber = func(playlistItem model.PlaylistItem, arrayIndex int) ItemWithNumber {
	return ItemWithNumber{
		Id:         playlistItem.Id,
		Title:      playlistItem.Title,
		Thumbnail:  playlistItem.Thumbnail,
		ItemNumber: arrayIndex + 1,
	}
}

var Index = template.FuncMap{
	"getItemWithNumber": getItemWithNumber,
}
