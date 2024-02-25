package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackematics/better-youtube-playlists/handler"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
	"github.com/jackematics/better-youtube-playlists/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestModalHiddenByDefault(t *testing.T) {
	state := page_data_repository.IndexState

	assert.Equal(t, true, state.ModalState.Hidden)
}

func TestModalOpens(t *testing.T) {
	req, err := http.NewRequest("GET", "/toggle-add-playlist-modal", nil)
	assert.Equal(t, nil, err)

	res_recorder := httptest.NewRecorder()

	handler.ToggleAddPlaylistModalHandler(res_recorder, req)

	assert.Equal(t, http.StatusOK, res_recorder.Code)

	body, err := io.ReadAll(res_recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("templates/add-playlist-modal.html", model.ModalModel{Hidden: false}), string(body)+"\r\n")
	assert.Equal(t, false, page_data_repository.IndexState.ModalState.Hidden)

	test_utils.ResetServerState()
}

func TestModalCloses(t *testing.T) {
	req, err := http.NewRequest("GET", "/toggle-add-playlist-modal", nil)
	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()
	handler.ToggleAddPlaylistModalHandler(recorder, req)
	recorder.Body.Reset()

	handler.ToggleAddPlaylistModalHandler(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("templates/add-playlist-modal.html", model.ModalModel{Hidden: true}), string(body)+"\r\n")
	assert.Equal(t, true, page_data_repository.IndexState.ModalState.Hidden)

	test_utils.ResetServerState()
}

func TestAddPlaylist(t *testing.T) {
	add_playlist_data := strings.NewReader("playlist_id=PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS")

	req, err := http.NewRequest("POST", "/add-playlist", add_playlist_data)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	handler.AddPlaylistHandler(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("templates/playlist-list-item.html", model.PlaylistModel{
		PlaylistId:    "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS",
		PlaylistTitle: "Better Youtube Playlists",
		ChannelOwner:  "Jack Rimmer",
	},
	), string(body)+"\r\n")

	test_utils.ResetServerState()
}
