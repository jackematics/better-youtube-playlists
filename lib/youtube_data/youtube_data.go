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

func FetchYoutubePlaylistMetadata(playlistId string) (*YoutubePlaylistMetadataResponse, *YoutubeDataError) {
	response, err := http.Get("https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&id=" + playlistId + "&key=" + config.YoutubeApiKey)

	if err != nil {
		log.Println("Error fetching youtube playlist metadata from https://youtube.googleapis.com/youtube/v3/playlists: " + err.Error())
		return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	rawBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println("Error reading body of youtube playlist metadata response: " + err.Error())
		return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	if response.StatusCode == http.StatusBadRequest {
		var youtubePlaylistResponseError YoutubePlaylistResponseError
		err = json.Unmarshal(rawBody, &youtubePlaylistResponseError)

		if err != nil {
			log.Println("Error decoding youtube metadata error response: ", err)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if youtubePlaylistResponseError.Error.Code == http.StatusBadRequest && youtubePlaylistResponseError.Error.Message == "API key not valid. Please pass a valid API key." {
			log.Println(youtubePlaylistResponseError.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if youtubePlaylistResponseError.Error.Code == http.StatusForbidden {
			log.Println(youtubePlaylistResponseError.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if youtubePlaylistResponseError.Error.Code >= http.StatusInternalServerError {
			log.Println("Issue retrieving data from Youtube Data API: " + youtubePlaylistResponseError.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}
	}

	var youtubePlaylistMetadataResponse YoutubePlaylistMetadataResponse
	err = json.Unmarshal(rawBody, &youtubePlaylistMetadataResponse)
	if err != nil {
		log.Println("Error decoding JSON response from Youtube api: ", err)
		return nil, &YoutubeDataError{Code: http.StatusFailedDependency, Message: "Internal server error"}
	}

	if len(youtubePlaylistMetadataResponse.Items) == 0 {
		log.Println("No playlist items returned for playlist id " + playlistId)

		return nil, &YoutubeDataError{Code: http.StatusBadRequest, Message: "Invalid playlist ID"}
	}

	return &youtubePlaylistMetadataResponse, nil
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

func fetchNextYoutubePlaylistItems(playlistId string, nextPageToken string) (*YoutubePlaylistItemsResponse, *YoutubeDataError) {
	url := "https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50" + "&playlistId=" + playlistId + "&key=" + config.YoutubeApiKey
	if nextPageToken != "" {
		url += "&pageToken=" + nextPageToken
	}

	youtubePlaylistItemsResponse, err := http.Get(url)

	if err != nil {
		log.Println("Error fetching youtube playlist items from https://youtube.googleapis.com/youtube/v3/playlistItems: " + err.Error())
		return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	responseData, err := io.ReadAll(youtubePlaylistItemsResponse.Body)

	if err != nil {
		log.Println("Error reading body of youtube playlist items response: " + err.Error())
		return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	}

	if youtubePlaylistItemsResponse.StatusCode == http.StatusBadRequest {
		var errorResponse YoutubePlaylistResponseError
		err = json.Unmarshal(responseData, &errorResponse)

		if err != nil {
			log.Println("Error decoding youtube playlist items error response: ", err)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if errorResponse.Error.Code == http.StatusBadRequest && errorResponse.Error.Message == "API key not valid. Please pass a valid API key." {
			log.Println(errorResponse.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if errorResponse.Error.Code == http.StatusForbidden {
			log.Println(errorResponse.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}

		if errorResponse.Error.Code >= http.StatusInternalServerError {
			log.Println("Issue retrieving data from Youtube Data API: " + errorResponse.Error.Message)
			return nil, &YoutubeDataError{Code: http.StatusInternalServerError, Message: "Internal server error"}
		}
	}

	var response_object YoutubePlaylistItemsResponse
	err = json.Unmarshal(responseData, &response_object)

	if err != nil {
		log.Println("Error decoding JSON response from Youtube api: ", err)
		return nil, &YoutubeDataError{Code: http.StatusFailedDependency, Message: "Error retrieving playlist data from Youtube"}
	}


	return &response_object, nil
}

func FetchYoutubePlaylistItems(playlistId string) (*YoutubePlaylistItemsResponse, *YoutubeDataError) {
	youtubePlaylistItemsResponse, youtubeDataErr := fetchNextYoutubePlaylistItems(playlistId, "")
	
	if youtubeDataErr != nil {
		return nil, youtubeDataErr
	}

	aggregatedItems := []Item{}
	aggregatedItems = append(aggregatedItems, youtubePlaylistItemsResponse.Items...)

	for youtubePlaylistItemsResponse.NextPageToken != "" {
		youtubePlaylistItemsResponse, youtubeDataErr = fetchNextYoutubePlaylistItems(playlistId, youtubePlaylistItemsResponse.NextPageToken)	

		if youtubeDataErr != nil {
			return nil, youtubeDataErr
		}

		aggregatedItems = append(aggregatedItems, youtubePlaylistItemsResponse.Items...)
	}

	youtubePlaylistItemsResponse.Items = aggregatedItems

	if len(youtubePlaylistItemsResponse.Items) == 0 {
		return nil, &YoutubeDataError{Code: http.StatusBadRequest, Message: "No playlist items returned"}
	}

	return youtubePlaylistItemsResponse, nil
}
