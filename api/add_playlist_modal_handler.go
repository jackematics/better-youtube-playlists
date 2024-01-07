package api

import (
	"net/http"

	"github.com/jackematics/better-youtube-playlists/repository"
	"github.com/jackematics/better-youtube-playlists/template"
)

func ToggleAddPlaylistModalHandler(w http.ResponseWriter, r *http.Request) {
	repository.ToggleAddPlaylistModal()

	modal := template.AddPlaylistModal(repository.GetPageState().ModalState)
	modal.Render(r.Context(), w)
}
