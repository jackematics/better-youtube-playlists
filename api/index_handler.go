package api

import (
	"net/http"

	"github.com/jackematics/better-youtube-playlists/repository"
	"github.com/jackematics/better-youtube-playlists/template"
)

func IndexRenderHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Index(*repository.GetPageState())
	index.Render(r.Context(), w)
}
