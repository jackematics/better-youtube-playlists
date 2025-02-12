package add_playlist

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jackematics/better-youtube-playlists/lib/youtube_data"
	"github.com/jackematics/better-youtube-playlists/model"
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

func AddPlaylistHandler(writer http.ResponseWriter, reader *http.Request) {
	playlistId := strings.TrimPrefix(reader.URL.Path, "/add-playlist/")

	if playlistId == "" {
		http.Error(writer, "Empty playlist ID", http.StatusBadRequest)
		return
	}

	youtubePlaylistMetadataResponse, youtubeDataError := youtube_data.FetchYoutubePlaylistMetadata(playlistId)

	if youtubeDataError != nil {
		http.Error(writer, youtubeDataError.Message, youtubeDataError.Code)
		return
	}

	playlistModel :=  model.Playlist{
		PlaylistId: playlistId,
		PlaylistTitle: youtubePlaylistMetadataResponse.Items[0].Snippet.Title,
		ChannelOwner: youtubePlaylistMetadataResponse.Items[0].Snippet.ChannelTitle,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(playlistModel)

	// youtube_playlist_items_response, youtube_err := youtube_data.FetchYoutubePlaylistItems(playlist_id)

	// if youtube_err != nil {
	// 	page_data.SetValidationMessage(youtube_err.Message)
	// 	http.Error(writer, youtube_err.Message, youtube_err.Code)
	// 	return
	// }

	// playlist_items := []model.PlaylistItem{}
	// for i := range youtube_playlist_items_response.Items {
	// 	response_item := youtube_playlist_items_response.Items[i]
	// 	item := model.PlaylistItem{
	// 		Id:           response_item.Snippet.ResourceId.VideoId,
	// 		Title:        response_item.Snippet.Title,
	// 		ThumbnailUrl: response_item.Snippet.Thumbnails.Default.Url,
	// 		Selected:     false,
	// 	}

	// 	playlist_items = append(playlist_items, item)
	// }

	// playlist_model := model.Playlist{
	// 	PlaylistId:    playlist_id,
	// 	PlaylistTitle: youtube_playlist_metadata_response.Items[0].Snippet.Title,
	// 	ChannelOwner:  youtube_playlist_metadata_response.Items[0].Snippet.ChannelTitle,
	// 	TotalVideos:   youtube_playlist_items_response.PageInfo.TotalResults,
	// 	Selected:      false,
	// 	PlaylistItems: playlist_items,
	// }

	// page_data.AddPlaylist(playlist_model)
	// page_data.ResetAddPlaylistValidation()

	// log.Println("Added playlist \"" + playlist_model.PlaylistTitle + "\" from playlist_id \"" + playlist_id + "\"")
	// tmpl := template.Must(template.ParseFiles("templates/playlist-list-item.html"))
	// tmpl.ExecuteTemplate(writer, "playlist-list-item", playlist_model)
}
