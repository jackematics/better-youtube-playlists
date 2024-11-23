package add_playlist

import (
	"log"
	"net/http"
	"text/template"

	"github.com/jackematics/better-youtube-playlists/lib/youtube_data"
	"github.com/jackematics/better-youtube-playlists/model"
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

	youtube_playlist_metadata_response, youtube_err := youtube_data.FetchYoutubePlaylistMetadata(playlist_id)

	if youtube_err != nil {
		page_data.SetValidationMessage(youtube_err.Message)
		http.Error(writer, youtube_err.Message, youtube_err.Code)
		return
	}

	youtube_playlist_items_response, youtube_err := youtube_data.FetchYoutubePlaylistItems(playlist_id)

	if youtube_err != nil {
		page_data.SetValidationMessage(youtube_err.Message)
		http.Error(writer, youtube_err.Message, youtube_err.Code)
		return
	}

	playlist_items := []model.PlaylistItem{}
	for i := range youtube_playlist_items_response.Items {
		response_item := youtube_playlist_items_response.Items[i]
		item := model.PlaylistItem{
			Id:           response_item.Snippet.ResourceId.VideoId,
			Title:        response_item.Snippet.Title,
			ThumbnailUrl: response_item.Snippet.Thumbnails.Default.Url,
			Selected:     false,
		}

		playlist_items = append(playlist_items, item)
	}

	playlist_model := model.Playlist{
		PlaylistId:    playlist_id,
		PlaylistTitle: youtube_playlist_metadata_response.Items[0].Snippet.Title,
		ChannelOwner:  youtube_playlist_metadata_response.Items[0].Snippet.ChannelTitle,
		TotalVideos:   youtube_playlist_items_response.PageInfo.TotalResults,
		Selected:      false,
		PlaylistItems: playlist_items,
	}

	page_data.AddPlaylist(playlist_model)
	page_data.ResetAddPlaylistValidation()

	log.Println("Added playlist \"" + playlist_model.PlaylistTitle + "\" from playlist_id \"" + playlist_id + "\"")
	tmpl := template.Must(template.ParseFiles("templates/playlist-list-item.html"))
	tmpl.ExecuteTemplate(writer, "playlist-list-item", playlist_model)
}
