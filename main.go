package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jackematics/better-youtube-playlists/template"
	"github.com/jackematics/better-youtube-playlists/model"

)

func main() {
	modal_state := model.ModalModel{Hidden: true}
	index_state := model.IndexModel{ModalState: modal_state}

	index := template.Index(index_state)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", templ.Handler(index))

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", niltempl generate --watch --proxy="http://localhost:8080" --cmd="go run .")
}
