package add_playlist

import (
	"log"
	"net/http"
	"text/template"

	"github.com/jackematics/better-youtube-playlists/helper/youtube_data"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
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

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type YoutubePlaylistMetadataResponseError struct {
	Error ErrorResponse `json:"error"`
}

func ToggleAddPlaylistModalWithValidationHandler(writer http.ResponseWriter, reader *http.Request) {
	if page_data.IndexState.ModalState.ValidationMessage == "" {
		page_data.ToggleAddPlaylistModal()
	}

	tmpl := template.Must(template.ParseFiles("templates/add-playlist-modal.html"))
	tmpl.ExecuteTemplate(writer, "add-playlist-modal", page_data.IndexState.ModalState)
}

func ToggleAddPlaylistModalHandler(writer http.ResponseWriter, reader *http.Request) {
	if !page_data.IndexState.ModalState.Hidden {
		page_data.ResetAddPlaylistValidation()
	}

	page_data.ToggleAddPlaylistModal()

	tmpl := template.Must(template.ParseFiles("templates/add-playlist-modal.html"))
	tmpl.ExecuteTemplate(writer, "add-playlist-modal", page_data.IndexState.ModalState)
}

func AddPlaylistHandler(writer http.ResponseWriter, reader *http.Request) {
	err := reader.ParseForm()
	if err != nil {
		http.Error(writer, "Failed to parse playlist id", http.StatusBadRequest)
	}

	playlist_id := reader.Form.Get("playlist_id")

	if playlist_id == "" {
		page_data.IndexState.ModalState.ValidationMessage = "Invalid playlist id"
		log.Println("Empty playlist_id")
		http.Error(writer, "Empty playlist_id", http.StatusBadRequest)
		return
	}

	_, duplicate_playlist_found := page_data.FindPlaylist(playlist_id)

	if duplicate_playlist_found {
		page_data.IndexState.ModalState.ValidationMessage = "Duplicate playlist id"
		log.Println("Duplicate playlist id: " + playlist_id)
		http.Error(writer, "Duplicate playlist id: "+playlist_id, http.StatusBadRequest)
		return
	}

	youtube_playlist, youtube_err := youtube_data.FetchYoutubeMetadata(playlist_id)

	if youtube_err != nil {
		page_data.SetValidationMessage(youtube_err.Message)
		http.Error(writer, youtube_err.Message, youtube_err.Code)
		return
	}

	page_data.AddPlaylist(*youtube_playlist)
	page_data.ResetAddPlaylistValidation()

	log.Println("Added playlist \"" + youtube_playlist.PlaylistTitle + "\" from playlist_id \"" + playlist_id + "\"")
	tmpl := template.Must(template.ParseFiles("templates/playlist-list-item.html"))
	tmpl.ExecuteTemplate(writer, "playlist-list-item", youtube_playlist)
}
