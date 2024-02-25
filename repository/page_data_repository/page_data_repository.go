package page_data_repository

import (
	"github.com/jackematics/better-youtube-playlists/model"
)

var IndexState = SetInitialState()

func SetInitialState() model.IndexModel {
	return model.IndexModel{
		ModalState: model.ModalModel{
			Hidden: true,
		},
		PlaylistState: []model.PlaylistModel{},
	}
}

func ToggleAddPlaylistModal() {
	IndexState.ModalState.Hidden = !IndexState.ModalState.Hidden
}

func AddPlaylist(playlist_model model.PlaylistModel) {
	IndexState.PlaylistState = append(IndexState.PlaylistState, playlist_model)
}
