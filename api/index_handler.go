package api

import (
	"net/http"

	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
	"github.com/jackematics/better-youtube-playlists/template"
)

func IndexRenderHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Index(page_data_repository.GetPageState())
	index.Render(r.Context(), w)
}
