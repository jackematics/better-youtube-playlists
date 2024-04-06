package func_map

import (
	"text/template"

	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

type ItemWithNumber struct {
	Id         string
	Title      string
	Thumbnail  model.Thumbnail
	ItemNumber int
}

func getItemWithNumber(playlistItem model.PlaylistItem, arrayIndex int) ItemWithNumber {
	return ItemWithNumber{
		Id:         playlistItem.Id,
		Title:      playlistItem.Title,
		Thumbnail:  playlistItem.Thumbnail,
		ItemNumber: arrayIndex + 1,
	}
}

func getSelected() *model.Playlist {
	selectedPlaylist, ok := page_data.FindSelected()

	if !ok {
		return &page_data.NilPlaylist
	}

	return selectedPlaylist
}

var PageFuncs = template.FuncMap{
	"getItemWithNumber": getItemWithNumber,
	"getSelected":       getSelected,
}
