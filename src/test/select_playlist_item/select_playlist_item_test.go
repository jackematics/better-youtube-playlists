package select_playlist_item

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackematics/better-youtube-playlists/handler/select_playlist_item"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
	_ "github.com/jackematics/better-youtube-playlists/test"
	"github.com/jackematics/better-youtube-playlists/test_utils"
	"github.com/stretchr/testify/assert"
)

func teardown() {
	test_utils.ResetServerState()
}

func TestHighlightSelectedPlaylistItem(t *testing.T) {
	playlist_list_state := []model.Playlist{
		{
			PlaylistId:    "test-id",
			PlaylistTitle: "Test Playlist",
			ChannelOwner:  "Test Owner",
			Selected:      true,
			PlaylistItems: []model.PlaylistItem{
				{
					Id:           "test-item-id-1",
					Title:        "Test Playlist Item 1",
					ThumbnailUrl: "test-thumbnail-url-1",
					Selected:     true,
				},
				{
					Id:           "test-item-id-2",
					Title:        "Test Playlist Item 2",
					ThumbnailUrl: "test-thumbnail-url-2",
					Selected:     false,
				},
			},
		},
	}

	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, playlist_list_state...)

	req, err := http.NewRequest("GET", "/highlight-selected-playlist-item/test-item-id-2", nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	select_playlist_item.HighlightSelectedPlaylistItemHandler(recorder, req)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("playlist-items", []string{"templates/playlist-items.html", "templates/playlist-item.html"}, page_data.IndexState.PlaylistListState[0]), string(body))
	assert.Equal(t, false, page_data.IndexState.PlaylistListState[0].PlaylistItems[0].Selected)
	assert.Equal(t, true, page_data.IndexState.PlaylistListState[0].PlaylistItems[1].Selected)

	teardown()
}

func TestPlaySelectedPlaylistItem(t *testing.T) {
	playlist_list_state := []model.Playlist{
		{
			PlaylistId:    "test-id",
			PlaylistTitle: "Test Playlist",
			ChannelOwner:  "Test Owner",
			Selected:      true,
			PlaylistItems: []model.PlaylistItem{
				{
					Id:           "test-item-id",
					Title:        "Test Playlist Item",
					ThumbnailUrl: "test-thumbnail-url",
					Selected:     true,
				},
			},
		},
	}

	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, playlist_list_state...)

	req, err := http.NewRequest("GET", "/play-selected-playlist-item/test-item-id", nil)

	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()

	select_playlist_item.PlaySelectedPlaylistItem(recorder, req)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("youtube-embed", []string{"templates/youtube-embed.html"}, page_data.IndexState.PlaylistListState[0].PlaylistItems[0]), string(body))

	teardown()
}

