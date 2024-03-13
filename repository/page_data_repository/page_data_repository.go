package page_data_repository

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
