package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/handler/add_playlist"
	"github.com/jackematics/better-youtube-playlists/handler/select_playlist"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func main() {
	tmpl, err := template.ParseFiles(
		"templates/index.html",
		"templates/add-playlist-modal.html",
		"templates/playlist-list-item.html",
		"templates/playlist-description.html",
	)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err != nil {
		log.Fatalf("could not init templates: %+v", err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		tmpl.ExecuteTemplate(writer, "index.html", page_data.IndexState)
	})

	http.HandleFunc("/toggle-add-playlist-modal", add_playlist.ToggleAddPlaylistModalHandler)
	http.HandleFunc("/toggle-add-playlist-modal-with-validation", add_playlist.ToggleAddPlaylistModalWithValidationHandler)
	http.HandleFunc("/add-playlist", add_playlist.AddPlaylistHandler)
	http.HandleFunc("/set-playlist-description", select_playlist.SetPlaylistDescriptionHandler)
	http.HandleFunc("/set-playlist-items", select_playlist.SetPlaylistItemsHandler)

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)
}
