package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	playlist "github.com/jackematics/better-youtube-playlists/handler/playlist"
)

//go:embed templates/*.html
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

func getContentType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".svg":
		return "image/svg+xml"
	default:
		return "text/plain"
	}
}

func main() {
	tmpl := template.Must(template.New("index").ParseFS(templateFS, "templates/*.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		filePath := request.URL.Path

		data, err := staticFS.ReadFile("static/" + filePath)
		if err != nil {
			http.NotFound(writer, request)
			return
		}

		contentType := getContentType(filePath)
		writer.Header().Set("Content-Type", contentType)

		writer.Write(data)
	})))

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		tmpl.ExecuteTemplate(writer, "index.html", nil)
	})

	http.HandleFunc("/playlist-metadata/", playlist.GetPlaylistMetadata)
	http.HandleFunc("/playlist-items/", playlist.GetPlaylistItems)

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", nil)	
}
