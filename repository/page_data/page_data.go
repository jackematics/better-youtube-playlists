package page_data

import (
	"github.com/jackematics/better-youtube-playlists/model"
)

var IndexState = InitialiseState()

func InitialiseState() model.IndexModel {
	return model.IndexModel{
		ModalState: model.ModalModel{
			Hidden:            true,
			ValidationMessage: "",
		},
		PlaylistState: []model.PlaylistModel{
			{
				PlaylistId:    "default-playlist-id",
				PlaylistTitle: "No Playlist Selected",
				ChannelOwner:  "",
			},
		},
	}
}

func ToggleAddPlaylistModal() {
	IndexState.ModalState.Hidden = !IndexState.ModalState.Hidden
}

func AddPlaylist(playlist_model model.PlaylistModel) {
	IndexState.PlaylistState = append(IndexState.PlaylistState, playlist_model)
}

func ResetAddPlaylistValidation() {
	IndexState.ModalState.ValidationMessage = ""
}

func SetValidationMessage(message string) {
	IndexState.ModalState.ValidationMessage = message
}

// returns true if the value is found and false otherwise
func FindPlaylist(playlist_id string) (*model.PlaylistModel, bool) {
	playlist_state := IndexState.PlaylistState
	for i := range playlist_state {
		if playlist_state[i].PlaylistId == playlist_id {
			return &playlist_state[i], true
		}
	}

	return nil, false
}
