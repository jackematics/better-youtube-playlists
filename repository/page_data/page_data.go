package page_data

import (
	"github.com/jackematics/better-youtube-playlists/model"
)

var IndexState = InitialiseState()

func InitialiseState() model.Index {
	return model.Index{
		ModalState: model.Modal{
			Hidden:            true,
			ValidationMessage: "",
		},
		PlaylistState: model.Playlist{
			Playlists: []model.PlaylistItem{
				{
					PlaylistId:    "default-playlist-id",
					PlaylistTitle: "No Playlist Selected",
					ChannelOwner:  "",
				},
			},
			SelectedPlaylistItemIndex: -1,
		},
	}
}

func ToggleAddPlaylistModal() {
	IndexState.ModalState.Hidden = !IndexState.ModalState.Hidden
}

func AddPlaylist(playlist_model model.PlaylistItem) {
	IndexState.PlaylistState.Playlists = append(IndexState.PlaylistState.Playlists, playlist_model)
}

func ResetAddPlaylistValidation() {
	IndexState.ModalState.ValidationMessage = ""
}

func SetValidationMessage(message string) {
	IndexState.ModalState.ValidationMessage = message
}

// returns true if the value is found and false otherwise
func FindPlaylist(playlist_id string) (*model.PlaylistItem, bool) {
	playlist_state := IndexState.PlaylistState.Playlists
	for i := range playlist_state {
		if playlist_state[i].PlaylistId == playlist_id {
			return &playlist_state[i], true
		}
	}

	return nil, false
}

func SetSelectedPage(playlist_id string) bool {
	playlists := IndexState.PlaylistState.Playlists
	for i := range playlists {
		if playlists[i].PlaylistId == playlist_id {
			IndexState.PlaylistState.SelectedPlaylistItemIndex = i
			return true
		}
	}

	return false
}
