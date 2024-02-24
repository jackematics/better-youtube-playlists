package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/handler"
	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
)

func main() {
	tmpl, err := template.ParseFiles(
		"templates/index.html",
		"templates/add-playlist-modal.html",
	)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err != nil {
		log.Fatalf("could not init templates: %+v", err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		tmpl.ExecuteTemplate(writer, "index.html", page_data_repository.IndexState)
	})

	http.HandleFunc("/toggle-add-playlist-modal", handler.ToggleAddPlaylistModalHandler)
	// http.HandleFunc("/add-playlist", handler.AddPlaylistHandler)

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)
}
