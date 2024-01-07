package main

import (
	"fmt"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/api"
	"github.com/jackematics/better-youtube-playlists/repository"
	"github.com/jackematics/better-youtube-playlists/template"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.IndexRenderHandler(w, r)
	})

	http.HandleFunc("/toggle-add-playlist-modal", api.ToggleAddPlaylistModalHandler)

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)
}

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	template.Index(repository.GetPageState()).Render(r.Context(), w)
}
