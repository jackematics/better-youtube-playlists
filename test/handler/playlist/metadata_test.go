package playlist

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	playlist "github.com/jackematics/better-youtube-playlists/handler/playlist"
	"github.com/jackematics/better-youtube-playlists/model"
	_ "github.com/jackematics/better-youtube-playlists/test"
	"github.com/jackematics/better-youtube-playlists/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestGetPlaylistMetadata(t *testing.T) {
	req, err := http.NewRequest("GET", "/playlist-metadata/" + test_utils.BETTER_YOUTUBE_PLAYLISTS_ID, nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	playlist.GetPlaylistMetadata(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedPlaylistMetadata := model.PlaylistMetadata{
		PlaylistId:    "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS",
		PlaylistTitle: "Better Youtube Playlists",
		ChannelOwner:  "Jack Rimmer",
	}

	rawBody, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)

	var body model.PlaylistMetadata
	json.Unmarshal(rawBody, &body)

	assert.Equal(t, expectedPlaylistMetadata.PlaylistId, body.PlaylistId)
	assert.Equal(t, expectedPlaylistMetadata.PlaylistTitle, body.PlaylistTitle)
	assert.Equal(t, expectedPlaylistMetadata.ChannelOwner, body.ChannelOwner)
}

func TestGetPlaylistMetadataFailsWithEmptyPlaylistId(t *testing.T) {
	req, err := http.NewRequest("GET", "/playlist-metadata/", nil)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	playlist.GetPlaylistMetadata(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Empty playlist ID\n", recorder.Body.String())
}

func TestGetPlaylistMetadataFailsWithInvalidPlaylistId(t *testing.T) {
	req, err := http.NewRequest("GET", "/playlist-metadata/test-invalid-id", nil)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	playlist.GetPlaylistMetadata(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Invalid playlist ID\n", recorder.Body.String())
}
