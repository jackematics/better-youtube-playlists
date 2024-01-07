package main

import (
	"fmt"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/api"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		api.IndexRenderHandler(writer, req)
	})

	http.HandleFunc("/toggle-add-playlist-modal", api.ToggleAddPlaylistModalHandler)

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)
}
