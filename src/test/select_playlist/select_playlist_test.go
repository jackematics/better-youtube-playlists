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

	req, err := http.NewRequest("GET", "/set-playlist-description/PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS", nil)

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
			PlaylistItems: []model.PlaylistItem{},
		},
		{
			PlaylistId:    "test-id-2",
			PlaylistTitle: "Test Playlist 2",
			ChannelOwner:  "Test Owner 2",
			Selected:      true,
			PlaylistItems: []model.PlaylistItem{},
		},
	}

	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, playlist_list_state...)

	req, err := http.NewRequest("GET", "/highlight-selected-playlist/test-id-2", nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	select_playlist.HighlightSelectedPlaylistHandler(recorder, req)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("playlist-list", []string{"templates/playlist-list.html", "templates/playlist-list-item.html"}, page_data.IndexState.PlaylistListState), string(body))
	assert.Equal(t, true, page_data.IndexState.PlaylistListState[1].Selected)

	teardown()
}

func TestPopulatePlaylistItems(t *testing.T) {
	test_playlist_id := "test-id"
	playlist_list_state := []model.Playlist{
		{
			PlaylistId:    test_playlist_id,
			PlaylistTitle: "Test Playlist",
			ChannelOwner:  "Test Owner",
			Selected:      true,
			PlaylistItems: []model.PlaylistItem{
				{
					Id:           "test-video-id",
					Title:        "Test Video Title",
					ThumbnailUrl: "https://test-thumbnail.com/id/default.jpg",
					Selected:     false,
				},
			},
		},
	}

	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, playlist_list_state...)

	req, err := http.NewRequest("GET", "/populate-playlist-items/"+test_playlist_id, nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	select_playlist.PopulatePlaylistItemsHandler(recorder, req)

	body, err := io.ReadAll(recorder.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("playlist-items", []string{"templates/playlist-items.html", "templates/playlist-item.html"}, page_data.IndexState.PlaylistListState[0]), string(body))

	teardown()
}
