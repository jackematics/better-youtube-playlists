package select_playlist

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackematics/better-youtube-playlists/handler/select_playlist"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
	_ "github.com/jackematics/better-youtube-playlists/test"
	"github.com/jackematics/better-youtube-playlists/test_utils"
	"github.com/stretchr/testify/assert"
)

func teardown() {
	test_utils.ResetServerState()
}

func TestSetPlaylistDescriptionHandler(t *testing.T) {
	playlist_item_data := model.PlaylistModel{
		PlaylistId:    "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS",
		PlaylistTitle: "Better Youtube Playlists",
		ChannelOwner:  "Jack Rimmer",
	}

	page_data_repository.IndexState.PlaylistState = append(page_data_repository.IndexState.PlaylistState, playlist_item_data)

	req, err := http.NewRequest("GET", "/set-playlist-description?playlist_id=PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS", nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	select_playlist.SetPlaylistDescriptionHandler(recorder, req)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("templates/playlist-description.html", playlist_item_data), string(body)+"\r\n")

	teardown()
}
