package api

import (
	"net/http"

	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
	"github.com/jackematics/better-youtube-playlists/template"
)

func ToggleAddPlaylistModalHandler(writer http.ResponseWriter, reader *http.Request) {
	page_data_repository.ToggleAddPlaylistModal()

	renderModal(writer, reader)
}

func renderModal(writer http.ResponseWriter, reader *http.Request) {
	modal := template.AddPlaylistModal(page_data_repository.GetPageState().ModalState)
	modal.Render(reader.Context(), writer)
}
