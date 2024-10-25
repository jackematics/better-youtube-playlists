package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/base64"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackematics/better-youtube-playlists/lib/func_map"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

type Response struct {
	StatusCode int `json:"statusCode"`
	Body string `json:"body"`
}

// var DBConnection *sql.DB
var UserId string

//go:embed templates/*.html
var templatesFs embed.FS

//go:embed static/*
var staticFs embed.FS

func main() {
	// http.HandleFunc("/toggle-add-playlist-modal", add_playlist.ToggleAddPlaylistModalHandler)
	// http.HandleFunc("/toggle-add-playlist-modal-with-validation", add_playlist.ToggleAddPlaylistModalWithValidationHandler)
	// http.HandleFunc("/add-playlist", add_playlist.AddPlaylistHandler)
	// http.HandleFunc("/set-playlist-description/", select_playlist.SetPlaylistDescriptionHandler)
	// http.HandleFunc("/highlight-selected-playlist/", select_playlist.HighlightSelectedPlaylistHandler)
	// http.HandleFunc("/populate-playlist-items/", select_playlist.PopulatePlaylistItemsHandler)
	// http.HandleFunc("/highlight-selected-playlist-item/", select_playlist_item.HighlightSelectedPlaylistItemHandler)
	// http.HandleFunc("/play-selected-playlist-item/", select_playlist_item.PlaySelectedPlaylistItem)

	
	fmt.Println("Lambda started")

	lambda.Start(lambdaHandler)
}


func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// tmpl := template.Must(template.New("index").Funcs(func_map.PageFuncs).ParseFS(templatesFs, "templates/*.html"))
	
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	// http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
	// 	tmpl.ExecuteTemplate(writer, "index.html", page_data.IndexState)
	// })

	if strings.HasPrefix(request.Path, "/static/") {
		return serveStaticFile(request.Path)
	}
	
	return serveHtmlTemplate()
}

func serveHtmlTemplate() (events.APIGatewayProxyResponse, error) {
	tmpl, err := template.New("index").Funcs(func_map.PageFuncs).ParseFS(templatesFs, "templates/*.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error"}, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, page_data.IndexState); err != nil {
		log.Printf("Error rendering template: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500, Body: "Internal Server Error"}, err
		}
	

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: buf.String(),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil
}

func serveStaticFile(path string) (events.APIGatewayProxyResponse, error) {
	relativePath := strings.TrimPrefix(path, "/")

	fileBytes, err := staticFs.ReadFile(relativePath)
	if err != nil {
		log.Printf("Error loading static file %s: %v", relativePath, err)
		return events.APIGatewayProxyResponse{ StatusCode: 404, Body: "File Not Found"}, nil
	}

	contentType := "text/plain"
	switch filepath.Ext(path) {
	case ".js":
		contentType = "application/javascript"
	case ".css":
		contentType = "text/css"
	case ".html":
		contentType = "text/html"
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".svg":
		contentType = "image/svg+xml"
	}
	
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": contentType,
		},
	}

	if strings.HasPrefix(contentType, "image/") {
		response.Body = base64.StdEncoding.EncodeToString(fileBytes)
		response.IsBase64Encoded = true
	} else {
		response.Body = string(fileBytes)
	}

	return response, nil
}