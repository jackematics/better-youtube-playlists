package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jackematics/better-youtube-playlists/templates"
)

func main() {
	index := templates.Index()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", templ.Handler(index))

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)
}
