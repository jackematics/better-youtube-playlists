package add_playlist

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackematics/better-youtube-playlists/handler/add_playlist"
	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
	_ "github.com/jackematics/better-youtube-playlists/test"
	"github.com/jackematics/better-youtube-playlists/test_utils"
	"github.com/stretchr/testify/assert"
)

func teardown() {
	test_utils.ResetServerState()
}

func TestModalHiddenByDefault(t *testing.T) {
	state := page_data.IndexState

	assert.Equal(t, true, state.ModalState.Hidden)
}

func TestModalOpens(t *testing.T) {
	req, err := http.NewRequest("GET", "/toggle-add-playlist-modal", nil)
	assert.Equal(t, nil, err)

	res_recorder := httptest.NewRecorder()

	add_playlist.ToggleAddPlaylistModalHandler(res_recorder, req)

	assert.Equal(t, http.StatusOK, res_recorder.Code)

	body, err := io.ReadAll(res_recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("add-playlist-modal", []string{"templates/add-playlist-modal.html"}, model.Modal{Hidden: false, ValidationMessage: ""}), string(body))
	assert.Equal(t, false, page_data.IndexState.ModalState.Hidden)

	teardown()
}

func TestModalCloses(t *testing.T) {
	req, err := http.NewRequest("GET", "/toggle-add-playlist-modal", nil)
	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()
	add_playlist.ToggleAddPlaylistModalHandler(recorder, req)
	recorder.Body.Reset()

	add_playlist.ToggleAddPlaylistModalHandler(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("add-playlist-modal", []string{"templates/add-playlist-modal.html"}, model.Modal{Hidden: true, ValidationMessage: ""}), string(body))
	assert.Equal(t, true, page_data.IndexState.ModalState.Hidden)

	teardown()
}

func TestModalStaysOpenWithValidationFailures(t *testing.T) {
	test_validation_message := "Test invalid"
	page_data.IndexState.ModalState.ValidationMessage = test_validation_message
	req, err := http.NewRequest("GET", "/toggle-add-playlist-modal-with-validation", nil)
	assert.Equal(t, nil, err)

	recorder := httptest.NewRecorder()
	add_playlist.ToggleAddPlaylistModalWithValidationHandler(recorder, req)
	recorder.Body.Reset()
	add_playlist.ToggleAddPlaylistModalHandler(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(
		t,
		test_utils.ParseTemplateToString("add-playlist-modal", []string{"templates/add-playlist-modal.html"}, model.Modal{Hidden: false, ValidationMessage: test_validation_message}),
		string(body),
	)
	assert.Equal(t, false, page_data.IndexState.ModalState.Hidden)
	assert.Equal(t, test_validation_message, page_data.IndexState.ModalState.ValidationMessage)

	teardown()
}

func TestAddPlaylist(t *testing.T) {
	add_playlist_data := strings.NewReader("playlist_id=PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS")

	req, err := http.NewRequest("POST", "/add_playlist", add_playlist_data)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	add_playlist.AddPlaylistHandler(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	playlist_item_data := model.Playlist{
		PlaylistId:    "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS",
		PlaylistTitle: "Better Youtube Playlists",
		ChannelOwner:  "Jack Rimmer",
		TotalVideos:   6,
		Selected:      false,
	}

	body, err := io.ReadAll(recorder.Body)
	assert.Equal(t, nil, err)
	assert.Equal(t, test_utils.ParseTemplateToString("playlist-list-item", []string{"templates/playlist-list-item.html"}, playlist_item_data), string(body))

	playlist_list_state := &page_data.IndexState.PlaylistListState

	assert.Equal(t, playlist_item_data.PlaylistId, (*playlist_list_state)[0].PlaylistId)
	assert.Equal(t, playlist_item_data.PlaylistTitle, (*playlist_list_state)[0].PlaylistTitle)
	assert.Equal(t, playlist_item_data.ChannelOwner, (*playlist_list_state)[0].ChannelOwner)

	assert.Equal(t, 6, len((*playlist_list_state)[0].PlaylistItems))
	assert.Equal(t, "snILjFUkk_A", ((*playlist_list_state)[0].PlaylistItems[0].Id))
	assert.Equal(t, "Depeche Mode - Never Let Me Down Again (Official Video) (Heard on Episode 1 of The Last Of Us)", ((*playlist_list_state)[0].PlaylistItems[0].Title))
	assert.Equal(t, "https://i.ytimg.com/vi/snILjFUkk_A/default.jpg", ((*playlist_list_state)[0].PlaylistItems[0].Thumbnail.Url))

	teardown()
}

func TestAddPlaylistFailsWithDuplicatePlaylist(t *testing.T) {
	test_playlist_id := "PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS"
	page_data.IndexState.PlaylistListState = append(page_data.IndexState.PlaylistListState, model.Playlist{
		PlaylistId:    test_playlist_id,
		PlaylistTitle: "",
		ChannelOwner:  "",
		Selected:      false,
	})

	add_playlist_data := strings.NewReader("playlist_id=" + test_playlist_id)

	req, err := http.NewRequest("POST", "/add_playlist", add_playlist_data)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	add_playlist.AddPlaylistHandler(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Duplicate playlist id: PLtcQcWdp-TodMQIlHfbpniiKVH9gHbiUS\n", string(recorder.Body.String()))
	assert.Equal(t, 1, len(page_data.IndexState.PlaylistListState))

	teardown()
}

func TestAddPlaylistFailsWithEmptyPlaylistId(t *testing.T) {
	add_playlist_data := strings.NewReader("playlist_id=")

	req, err := http.NewRequest("POST", "/add_playlist", add_playlist_data)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	add_playlist.AddPlaylistHandler(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Empty playlist_id\n", recorder.Body.String())

	teardown()
}

func TestAddPlaylistFailsWithInvalidPlaylistId(t *testing.T) {
	test_invalid_id := "test-invalid-id"
	add_playlist_data := strings.NewReader("playlist_id=" + test_invalid_id)

	req, err := http.NewRequest("POST", "/add_playlist", add_playlist_data)

	assert.Equal(t, nil, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	add_playlist.AddPlaylistHandler(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Invalid playlist id\n", recorder.Body.String())

	teardown()
}
