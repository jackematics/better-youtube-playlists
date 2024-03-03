package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/jackematics/better-youtube-playlists/config"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
)

type Snippet struct {
	Title        string `json:"title"`
	ChannelTitle string `json:"channelTitle"`
}

type MetadataItem struct {
	Id      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type YoutubePlaylistMetadataResponse struct {
	Items []MetadataItem `json:"items"`
}

func ToggleAddPlaylistModalWithValidationHandler(writer http.ResponseWriter, reader *http.Request) {
	if page_data_repository.IndexState.ModalState.ValidationMessage == "" {
		page_data_repository.ToggleAddPlaylistModal()
	}

	tmpl := template.Must(template.ParseFiles("templates/add-playlist-modal.html"))
	tmpl.ExecuteTemplate(writer, "add-playlist-modal", page_data_repository.IndexState.ModalState)
}

func ToggleAddPlaylistModalHandler(writer http.ResponseWriter, reader *http.Request) {
	if !page_data_repository.IndexState.ModalState.Hidden {
		page_data_repository.ResetAddPlaylistValidation()
	}

	page_data_repository.ToggleAddPlaylistModal()

	tmpl := template.Must(template.ParseFiles("templates/add-playlist-modal.html"))
	tmpl.ExecuteTemplate(writer, "add-playlist-modal", page_data_repository.IndexState.ModalState)
}

func AddPlaylistHandler(writer http.ResponseWriter, reader *http.Request) {
	err := reader.ParseForm()
	if err != nil {
		http.Error(writer, "Failed to parse playlist id", http.StatusBadRequest)
	}

	playlist_id := reader.Form.Get("playlist_id")

	if playlist_id == "" {
		page_data_repository.IndexState.ModalState.ValidationMessage = "Invalid playlist id"
		log.Println("Empty playlist_id")
		http.Error(writer, "Empty playlist_id", http.StatusBadRequest)
		return
	}

	for i := range page_data_repository.IndexState.PlaylistState {
		if page_data_repository.IndexState.PlaylistState[i].PlaylistId == playlist_id {
			page_data_repository.IndexState.ModalState.ValidationMessage = "Duplicate playlist id"
			log.Println("Duplicate playlist id: " + playlist_id)
			http.Error(writer, "Duplicate playlist id: "+playlist_id, http.StatusBadRequest)
			return
		}
	}

	youtube_playlist_metadata_response, _ := http.Get("https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&id=" + playlist_id + "&key=" + config.Config.YoutubeApiKey)

	response_data, _ := io.ReadAll(youtube_playlist_metadata_response.Body)

	var response_object YoutubePlaylistMetadataResponse
	err = json.Unmarshal(response_data, &response_object)
	if err != nil {
		log.Println("Error decoding JSON response from Youtube api: ", err)
		http.Error(writer, "Failed Dependency", http.StatusFailedDependency)
		return
	}

	if len(response_object.Items) == 0 {
		log.Println("No playlist items returned for playlist id " + playlist_id)
		page_data_repository.IndexState.ModalState.ValidationMessage = "Invalid playlist id"
		http.Error(writer, "Invalid playlist_id: "+playlist_id, http.StatusBadRequest)

		return
	}

	playlist_model := model.PlaylistModel{
		PlaylistId:    playlist_id,
		PlaylistTitle: response_object.Items[0].Snippet.Title,
		ChannelOwner:  response_object.Items[0].Snippet.ChannelTitle,
	}

	page_data_repository.AddPlaylist(playlist_model)
	page_data_repository.ResetAddPlaylistValidation()

	log.Println("Added playlist \"" + playlist_model.PlaylistTitle + "\" from playlist_id \"" + playlist_id + "\"")
	tmpl := template.Must(template.ParseFiles("templates/playlist-list-item.html"))
	tmpl.ExecuteTemplate(writer, "playlist-list-item", playlist_model)
}
