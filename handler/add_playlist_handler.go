package handler

import (
	"net/http"
	"text/template"

	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
)

// type Snippet struct {
// 	Title        string `json:"title"`
// 	ChannelTitle string `json:"channelTitle"`
// }

// type MetadataItem struct {
// 	Id      string  `json:"id"`
// 	Snippet Snippet `json:"snippet"`
// }

// type YoutubePlaylistMetadataResponse struct {
// 	Items []MetadataItem `json:"items"`
// }

func ToggleAddPlaylistModalHandler(writer http.ResponseWriter, reader *http.Request) {
	page_data_repository.ToggleAddPlaylistModal()

	tmpl := template.Must(template.ParseFiles("templates/add-playlist-modal.html"))
	tmpl.ExecuteTemplate(writer, "add-playlist-modal", page_data_repository.IndexState.ModalState)
}

// func AddPlaylistHandler(writer http.ResponseWriter, reader *http.Request) {
// reader.ParseForm()

// playlist_id := reader.FormValue("playlist_id")

// youtube_playlist_metadata_response, _ := http.Get("https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&id=" + playlist_id + "&key=" + youtube_api_key)
// response_data, _ := io.ReadAll(youtube_playlist_metadata_response.Body)

// var response_object YoutubePlaylistMetadataResponse
// err := json.Unmarshal(response_data, &response_object)
// if err != nil {
// 	http.Error(writer, "Error decoding JSON response", http.StatusInternalServerError)
// 	return
// }

// playlist_model := model.PlaylistModel{
// 	PlaylistId:    playlist_id,
// 	PlaylistTitle: response_object.Items[0].Snippet.Title,
// 	ChannelOwner:  response_object.Items[0].Snippet.ChannelTitle,
// }

// playlist_state := page_data_repository.AddPlaylist(playlist_model)

// playlist_list := template.PlaylistList(playlist_state)
// playlist_list.Render(reader.Context(), writer)
// }
