package youtube_data

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/config"
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

type YoutubePlaylistResponseError struct {
	Error ErrorResponse `json:"error"`
}

type YoutubeDataError struct {
	Code    int
	Message string
}

func FetchYoutubePlaylistMetadata(playlist_id string) (*YoutubePlaylistMetadataResponse, *YoutubeDataError) {
	youtube_playlist_metadata_response, err := http.Get("https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&id=" + playlist_id + "&key=" + config.Config.YoutubeApiKey)

	if err != nil {
		log.Println("Error fetching youtube playlist metadata from https://youtube.googleapis.com/youtube/v3/playlists: " + err.Error())
		return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	response_data, err := io.ReadAll(youtube_playlist_metadata_response.Body)

	if err != nil {
		log.Println("Error reading body of youtube playlist metadata response: " + err.Error())
		return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	if youtube_playlist_metadata_response.StatusCode == http.StatusBadRequest {
		var error_response YoutubePlaylistResponseError
		err = json.Unmarshal(response_data, &error_response)

		if err != nil {
			log.Println("Error decoding youtube metadata error response: ", err)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if error_response.Error.Code == http.StatusBadRequest && error_response.Error.Message == "API key not valid. Please pass a valid API key." {
			log.Println(error_response.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if error_response.Error.Code == http.StatusForbidden {
			log.Println(error_response.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if error_response.Error.Code >= http.StatusInternalServerError {
			log.Println("Issue retrieving data from Youtube Data API: " + error_response.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}
	}

	var response_object YoutubePlaylistMetadataResponse
	err = json.Unmarshal(response_data, &response_object)
	if err != nil {
		log.Println("Error decoding JSON response from Youtube api: ", err)
		return nil, &YoutubeDataError{Code: http.StatusFailedDependency, Message: "Error retrieving playlist data from Youtube"}
	}

	if len(response_object.Items) == 0 {
		log.Println("No playlist items returned for playlist id " + playlist_id)

		return nil, &YoutubeDataError{Code: http.StatusBadRequest, Message: "Invalid playlist id"}
	}

	return &response_object, nil
}

type Thumbnail struct {
	Url string `json:"url"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
}

type ResourceId struct {
	VideoId string `json:"videoId"`
}

type ItemSnippet struct {
	Title      string     `json:"title"`
	Thumbnails Thumbnails `json:"thumbnails"`
	ResourceId ResourceId `json:"resourceId"`
}

type Item struct {
	Snippet ItemSnippet `json:"snippet"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type YoutubePlaylistItemsResponse struct {
	NextPageToken string   `json:"nextPageToken"`
	Items         []Item   `json:"items"`
	PageInfo      PageInfo `json:"pageInfo"`
}

func FetchYoutubePlaylistItems(playlist_id string) (*YoutubePlaylistItemsResponse, *YoutubeDataError) {
	youtube_playlist_items_response, err := http.Get("https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50" + "&playlistId=" + playlist_id + "&key=" + config.Config.YoutubeApiKey)

	if err != nil {
		log.Println("Error fetching youtube playlist items from https://youtube.googleapis.com/youtube/v3/playlistItems: " + err.Error())
	}

	response_data, err := io.ReadAll(youtube_playlist_items_response.Body)

	if err != nil {
		log.Println("Error reading body of youtube playlist items response: " + err.Error())
		return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	if youtube_playlist_items_response.StatusCode == http.StatusBadRequest {
		var error_response YoutubePlaylistResponseError
		err = json.Unmarshal(response_data, &error_response)

		if err != nil {
			log.Println("Error decoding youtube playlist items error response: ", err)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if error_response.Error.Code == http.StatusBadRequest && error_response.Error.Message == "API key not valid. Please pass a valid API key." {
			log.Println(error_response.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if error_response.Error.Code == http.StatusForbidden {
			log.Println(error_response.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if error_response.Error.Code >= http.StatusInternalServerError {
			log.Println("Issue retrieving data from Youtube Data API: " + error_response.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}
	}

	var response_object YoutubePlaylistItemsResponse
	err = json.Unmarshal(response_data, &response_object)

	if err != nil {
		log.Println("Error decoding JSON response from Youtube api: ", err)
		return nil, &YoutubeDataError{Code: http.StatusFailedDependency, Message: "Error retrieving playlist data from Youtube"}
	}

	if len(response_object.Items) == 0 {
		log.Println("No playlist items returned for playlist id " + playlist_id)

		return nil, &YoutubeDataError{Code: http.StatusBadRequest, Message: "Invalid playlist id"}
	}

	return &response_object, nil
}
