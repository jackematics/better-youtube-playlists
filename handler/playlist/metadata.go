package playlist

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

func GetPlaylistMetadata(writer http.ResponseWriter, reader *http.Request) {
	playlistId := strings.TrimPrefix(reader.URL.Path, "/playlist-metadata/")

	if playlistId == "" {
		http.Error(writer, "Empty playlist ID", http.StatusBadRequest)
		return
	}

	youtubePlaylistMetadataResponse, youtubeDataError := youtube_data.FetchYoutubePlaylistMetadata(playlistId)

	if youtubeDataError != nil {
		http.Error(writer, youtubeDataError.Message, youtubeDataError.Code)
		return
	}

	playlistModel :=  model.PlaylistMetadata{
		PlaylistId: playlistId,
		PlaylistTitle: youtubePlaylistMetadataResponse.Items[0].Snippet.Title,
		ChannelOwner: youtubePlaylistMetadataResponse.Items[0].Snippet.ChannelTitle,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(playlistModel)
}
