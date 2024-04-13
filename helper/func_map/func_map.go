package func_map

import (
	"text/template"

	"github.com/jackematics/better-youtube-playlists/model"
)

type ItemWithNumber struct {
	Id           string
	Title        string
	ThumbnailUrl string
	ItemNumber   int
	Selected     bool
}

func getItemWithNumber(playlistItem model.PlaylistItem, arrayIndex int) ItemWithNumber {
	return ItemWithNumber{
		Id:           playlistItem.Id,
		Title:        playlistItem.Title,
		ThumbnailUrl: playlistItem.ThumbnailUrl,
		ItemNumber:   arrayIndex + 1,
		Selected:     playlistItem.Selected,
	}
}

var nilPlaylist = model.Playlist{
	PlaylistId:    "default-playlist-id",
	PlaylistTitle: "No Playlist Selected",
	ChannelOwner:  "",
	TotalVideos:   0,
	Selected:      false,
	PlaylistItems: []model.PlaylistItem{},
}

func getSelectedPlaylist(playlist_list []model.Playlist) model.Playlist {
	for i := range playlist_list {
		playlist := (playlist_list)[i]
		if (playlist_list)[i].Selected {
			return playlist
		}
	}

	return nilPlaylist
}

var nilPlaylistItem = model.PlaylistItem{
	Id:           "",
	Title:        "",
	ThumbnailUrl: "",
	Selected:     false,
}

func getSelectedPlaylistItem(playlist model.Playlist) *model.PlaylistItem {
	for _, item := range playlist.PlaylistItems {
		if item.Selected {
			return &item
		}
	}

	return &nilPlaylistItem
}

var PageFuncs = template.FuncMap{
	"getItemWithNumber":       getItemWithNumber,
	"getSelectedPlaylist":     getSelectedPlaylist,
	"getSelectedPlaylistItem": getSelectedPlaylistItem,
}
