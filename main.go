package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jackematics/better-youtube-playlists/template"
)

func main() {
	modal_state := ModalState{hidden: true}
	index_state := IndexState{modal_state: modal_state}

	index := template.Index(index_state)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", templ.Handler(index))

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)
}
