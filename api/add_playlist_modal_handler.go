package api

import (
	"net/http"

	"github.com/jackematics/better-youtube-playlists/repository"
)

func ToggleAddPlaylistModalHandler(w http.ResponseWriter, r *http.Request) {
	repository.ToggleAddPlaylistModal()

	IndexRenderHandler(w, r)
}
