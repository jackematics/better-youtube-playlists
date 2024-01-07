package repository

import "github.com/jackematics/better-youtube-playlists/model"

var modal_state = model.ModalModel{
	Hidden: true,
}

var IndexState = model.IndexModel{
	ModalState: modal_state,
}

func GetPageState() *model.IndexModel {
	return &IndexState
}
