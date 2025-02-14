package playlist

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackematics/better-youtube-playlists/handler/playlist"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestGetPlaylistItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/playlist-items/" + test_utils.BETTER_YOUTUBE_PLAYLISTS_ID, nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	playlist.GetPlaylistItems(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	rawBody, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)

	var body model.Playlist
	json.Unmarshal(rawBody, &body)
	
	assert.Equal(t, 6, body.TotalVideos)
	assert.Equal(t, "snILjFUkk_A", (body.Items[0].Id))
	assert.Equal(t, "Depeche Mode - Never Let Me Down Again (Remastered)", (body.Items[0].Title))
	assert.Equal(t, "https://i.ytimg.com/vi/snILjFUkk_A/default.jpg", (body.Items[0].ThumbnailUrl))
}

func TestGetPlaylistItemsFailsWithEmptyPlaylistId(t *testing.T) {
	req, err := http.NewRequest("GET", "/playlist-items/", nil)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	playlist.GetPlaylistItems(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Empty playlist ID\n", recorder.Body.String())
}