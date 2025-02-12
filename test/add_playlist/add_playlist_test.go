package add_playlist

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	addPlaylist "github.com/jackematics/better-youtube-playlists/handler/add_playlist"
	"github.com/jackematics/better-youtube-playlists/model"
	_ "github.com/jackematics/better-youtube-playlists/test"
	"github.com/stretchr/testify/assert"
)

func TestAddPlaylist(t *testing.T) {
	req, err := http.NewRequest("GET", "/add-playlist/PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS", nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	addPlaylist.AddPlaylistHandler(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedPlaylistMetadata := model.Playlist{
		PlaylistId:    "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS",
		PlaylistTitle: "Better Youtube Playlists",
		ChannelOwner:  "Jack Rimmer",
	}

	rawBody, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)

	var body model.Playlist
	json.Unmarshal(rawBody, &body)

	assert.Equal(t, expectedPlaylistMetadata.PlaylistId, body.PlaylistId)
	assert.Equal(t, expectedPlaylistMetadata.PlaylistTitle, body.PlaylistTitle)
	assert.Equal(t, expectedPlaylistMetadata.ChannelOwner, body.ChannelOwner)
}

// func TestAddPlaylistFailsWithDuplicatePlaylist(t *testing.T) {
// 	test_playlist_id := "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS"
// 	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, model.Playlist{
// 		PlaylistId:    test_playlist_id,
// 		PlaylistTitle: "",
// 		ChannelOwner:  "",
// 		Selected:      false,
// 	})

// 	add_playlist_data := strings.NewReader("playlist_id=" + test_playlist_id)

// 	req, err := http.NewRequest("POST", "/add_playlist", add_playlist_data)

// 	assert.Equal(t, nil, err)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	recorder := httptest.NewRecorder()

// 	add_playlist.AddPlaylistHandler(recorder, req)

// 	assert.Equal(t, http.StatusBadRequest, recorder.Code)
// 	assert.Equal(t, "Duplicate playlist id: PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS\n", string(recorder.Body.String()))
// 	assert.Equal(t, 1, len(page_data.IndexState.PlaylistListState))

// 	teardown()
// }

// func TestAddPlaylistFailsWithEmptyPlaylistId(t *testing.T) {
// 	add_playlist_data := strings.NewReader("playlist_id=")

// 	req, err := http.NewRequest("GET", "/add-playlist/playlist_id", add_playlist_data)

// 	assert.Equal(t, nil, err)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	recorder := httptest.NewRecorder()

// 	add_playlist.AddPlaylistHandler(recorder, req)

// 	assert.Equal(t, http.StatusBadRequest, recorder.Code)
// 	assert.Equal(t, "Empty playlist_id\n", recorder.Body.String())

// 	teardown()
// }

// func TestAddPlaylistFailsWithInvalidPlaylistId(t *testing.T) {
// 	test_invalid_id := "test-invalid-id"
// 	add_playlist_data := strings.NewReader("playlist_id=" + test_invalid_id)

// 	req, err := http.NewRequest("POST", "/add_playlist", add_playlist_data)

// 	assert.Equal(t, nil, err)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	recorder := httptest.NewRecorder()

// 	add_playlist.AddPlaylistHandler(recorder, req)

// 	assert.Equal(t, http.StatusBadRequest, recorder.Code)
// 	assert.Equal(t, "Invalid playlist id\n", recorder.Body.String())

// 	teardown()
// }
