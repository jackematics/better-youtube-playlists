package select_playlist

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackematics/better-youtube-playlists/handler/select_playlist"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
	_ "github.com/jackematics/better-youtube-playlists/test"
	"github.com/jackematics/better-youtube-playlists/test_utils"
	"github.com/stretchr/testify/assert"
)

func teardown() {
	test_utils.ResetServerState()
}

func TestSetPlaylistDescriptionHandler(t *testing.T) {
	playlist_item_data := model.Playlist{
		PlaylistId:    "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS",
		PlaylistTitle: "Better Youtube Playlists",
		ChannelOwner:  "Jack Rimmer",
		Selected:      false,
	}

	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, playlist_item_data)

	req, err := http.NewRequest("GET", "/set-playlist-description?playlist_id=PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS", nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	select_playlist.SetPlaylistDescriptionHandler(recorder, req)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("playlist-description", []string{"templates/playlist-description.html"}, playlist_item_data), string(body))

	teardown()
}

func TestHighlightSelectedPlaylist(t *testing.T) {
	playlist_list_state := []model.Playlist{
		{
			PlaylistId:    "test-id-1",
			PlaylistTitle: "Test Playlist 1",
			ChannelOwner:  "Test Owner 1",
			Selected:      false,
		},
		{
			PlaylistId:    "test-id-2",
			PlaylistTitle: "Test Playlist 2",
			ChannelOwner:  "Test Owner 2",
			Selected:      true,
		},
	}

	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, playlist_list_state...)

	req, err := http.NewRequest("GET", "/highlight-selected-playlist/test-id-2", nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	select_playlist.HighlightSelectedPlaylist(recorder, req)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("playlist-list", []string{"templates/playlist-list.html", "templates/playlist-list-item.html"}, page_data.IndexState.PlaylistListState), string(body))

	teardown()
}
