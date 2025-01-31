package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/handler/add_playlist"
	"github.com/jackematics/better-youtube-playlists/handler/select_playlist"
	"github.com/jackematics/better-youtube-playlists/handler/select_playlist_item"
	"github.com/jackematics/better-youtube-playlists/lib/func_map"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

//go:embed templates/*.html
var templateFS embed.FS

func main() {
	tmpl := template.Must(template.New("index").Funcs(func_map.PageFuncs).ParseFS(templateFS, "templates/*.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		tmpl.ExecuteTemplate(writer, "index.html", page_data.IndexState)
	})

	http.HandleFunc("/toggle-add-playlist-modal", add_playlist.ToggleAddPlaylistModalHandler)
	http.HandleFunc("/toggle-add-playlist-modal-with-validation", add_playlist.ToggleAddPlaylistModalWithValidationHandler)
	http.HandleFunc("/add-playlist", add_playlist.AddPlaylistHandler)
	http.HandleFunc("/set-playlist-description/", select_playlist.SetPlaylistDescriptionHandler)
	http.HandleFunc("/highlight-selected-playlist/", select_playlist.HighlightSelectedPlaylistHandler)
	http.HandleFunc("/populate-playlist-items/", select_playlist.PopulatePlaylistItemsHandler)
	http.HandleFunc("/highlight-selected-playlist-item/", select_playlist_item.HighlightSelectedPlaylistItemHandler)
	http.HandleFunc("/play-selected-playlist-item/", select_playlist_item.PlaySelectedPlaylistItem)

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)	
}
