package repository

import (
	"github.com/jackematics/better-youtube-playlists/model"
)

var IndexState = model.IndexModel{
	ModalState: model.ModalModel{
		Hidden: true,
	},
}

func GetPageState() *model.IndexModel {
	return &IndexState
}

func ToggleAddPlaylistModal() {
	IndexState.ModalState.Hidden = !IndexState.ModalState.Hidden
}
