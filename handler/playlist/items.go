package playlist

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/jackematics/better-youtube-playlists/lib/youtube_data"
	"github.com/jackematics/better-youtube-playlists/model"
)

func GetPlaylistItems(writer http.ResponseWriter, reader *http.Request) {
	playlistId := strings.TrimPrefix(reader.URL.Path, "/playlist-items/")

	if playlistId == "" {
		http.Error(writer, "Empty playlist ID", http.StatusBadRequest)
		return
	}

	playlistIdRegex := regexp.MustCompile("^PL[A-Za-z0-9_-]+$")

	if !playlistIdRegex.MatchString(playlistId) {
		http.Error(writer, "Invalid playlist ID", http.StatusBadRequest)
		return
	}

	youtubePlaylistItemsResponse, youtubeDataErr := youtube_data.FetchYoutubePlaylistItems(playlistId)

	if youtubeDataErr != nil {
		http.Error(writer, youtubeDataErr.Message, youtubeDataErr.Code)
		return
	}

	playlist := model.Playlist{}
	playlist.TotalVideos = youtubePlaylistItemsResponse.PageInfo.TotalResults

	for i := range youtubePlaylistItemsResponse.Items {
		responseItem := youtubePlaylistItemsResponse.Items[i]
		item := model.Item{
			Id:           responseItem.Snippet.ResourceId.VideoId,
			Title:        responseItem.Snippet.Title,
			ThumbnailUrl: responseItem.Snippet.Thumbnails.Default.Url,
		}

		playlist.Items = append(playlist.Items, item)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(playlist)
}